// Copyright © 2019 Alexey Stolpovskikh <stolpovskikh@hotmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/alexstov/sling/conf"
	"github.com/alexstov/sling/emul"
	"github.com/alexstov/sling/net"
	"github.com/alexstov/sling/sio"
	"github.com/alexstov/sling/throt"
	"github.com/cloudflare/cfssl/log"
	"github.com/rcrowley/go-metrics"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cmdSend = &cobra.Command{
	Use:   "send [flags]",
	Short: "Send requests to the endpoint",
	Long: `Send requests from a file or directory to the endpoint.

Examples:
# Send myfile.dat from /tmp/data directory. Use default send parameters from SLINGCONFIG.
sling request send -f myfile.dat -d /tmp/data`,
	Run: sendRun,
}

// Send command flags
var flagmapper Flagmapper

// Send command metrics
var reg metrics.Registry

// SessionID Send command session ID
var SessionID string

func init() {
	var err error

	// Generate unique session ID.
	t := time.Now()
	SessionID = fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond())

	// Create send command flags.
	if flagmapper, err = NewCmdFlags(cmdSend, CmdSend, sconf); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create send flags.")
	}

	// Enable event for active endpoint index in case it's passed explicitly.
	flagmapper.GetFlagmap()[Endpoint].AddEvent(Events[ActiveEndpointSet], flagmapper.SetActiveEndpoint)

	// Add send subcommand to request command.
	RequestCmd.AddCommand(cmdSend)

	// Create new send command metrics registry.
	reg = metrics.NewRegistry()

	logger.Out(logrus.TraceLevel, nil, "Send command initialized.")
}

func sendRun(cmd *cobra.Command, args []string) {
	var err error
	var sendArgs emul.SendArgs
	var em *emul.Emul

	// Create background context for synchronization.
	ctx := context.Background()

	// Set explicit command flags to override config defaults.
	RootFlags.SetExplicit()
	flagmapper.SetExplicit()

	// Resolve send command arguments, filepath, etc. from the command flags.
	if err = flagmapper.ResolveSendArgs(&sendArgs); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot resolve send arguments.")
	}

	// Create new emulator.
	if em, err = NewSendEmul(&sendArgs); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create emul.")
	}

	// Make channel large enough to store all requests.
	in := make(chan interface{}, sendArgs.Repeat)

	switch sendArgs.SendType {
	case emul.SingleReq:
		sendArgs.ReqID = 1

		// Send a single request.
		err = em.Dispatcher.SendReq(ctx, sendArgs.Data, &sendArgs)

	case emul.RepeatReq, emul.MultiReq:
		var wg sync.WaitGroup
		wg.Add(2)

		// Start gouroutine to prepare all requests
		go prepareRequests(in, &sendArgs, &wg)

		// Start gouroutine to send all requests.
		go em.Dispatcher.MultiSend(ctx, in, &sendArgs, &wg)

		// Wait for all requests to prepare and send.
		wg.Wait()
	}

	// Log the histogramm.
	flagmap := flagmapper.GetFlagmap()
	if flag, ok := flagmap[LogHis]; ok {
		logHis := flag.Value.(*BoolVal).Value
		if logHis {
			logHistoCh := make(chan interface{})
			go metrics.LogScaledOnCue(reg, logHistoCh, time.Millisecond, logger.GetLogger())
			logHistoCh <- 1
			defer close(logHistoCh)
		}
	}

	// Print the histogram to console.
	if flag, ok := flagmap[ConHis]; ok {
		conHis := flag.Value.(*BoolVal).Value
		if conHis {
			conHistoCh := make(chan interface{})
			go metrics.LogScaledOnCue(reg, conHistoCh, time.Millisecond, em.Consoler.GetLogger())
			conHistoCh <- 1
			defer close(conHistoCh)
		}
	}

	// Output commanad results.
	if err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "Send command execution failed.")
	} else {
		logger.Out(logrus.InfoLevel, nil, "Send command executed.")
	}
}

// NewSendEmul creates new send emulator.
func NewSendEmul(sendArgs *emul.SendArgs) (em *emul.Emul, err error) {
	var client net.Client
	var filer sio.Filer
	var limiter throt.Limiter

	// Create emul filer.
	if filer, err = sio.NewFiler(logger); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create filer.")
	}

	// Create emul limiter.
	if limiter, err = throt.NewMultiLimiter(
		&throt.MultiLimitArgs{RateSec: sendArgs.RateSec,
			RateMin: sendArgs.RateMin,
			CxnNum:  sendArgs.CxnNum}); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create limiter.")
	}

	switch sendArgs.CltType {
	case conf.TCP:
		if client, err = net.NewTCPClient(logger, filer); err != nil {
			logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create TCP client.")
		}
	case conf.HTTPPost:
		if client, err = net.NewHTTPClient(logger, filer); err != nil {
			logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create HTTP client.")
		}
	}
	logger.Out(logrus.InfoLevel, logrus.Fields{"ClientType": sendArgs.CltType}, "Set client type.")

	// Create histogram.
	msamp := metrics.NewUniformSample(1028) // metrics.NewExpDecaySample(1028, 0.015)
	histo := metrics.NewHistogram(msamp)
	reg.Register("Client", histo)
	if em, err = emul.NewEmul(client, filer, Con, limiter, logger, histo); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create emul.")
	}

	// Create SaveReq directory.
	if sendArgs.SaveReq {
		if _, err := os.Stat(sendArgs.SaveReqDir); os.IsNotExist(err) {
			os.Mkdir(sendArgs.SaveReqDir, 0755)
		}
	}

	// Create SaveRes directory.
	if sendArgs.SaveRes {
		if _, err := os.Stat(sendArgs.SaveResDir); os.IsNotExist(err) {
			os.Mkdir(sendArgs.SaveResDir, 0755)
		}
	}

	return em, err
}

// prepareRequests prepares requests to send.
func prepareRequests(out chan<- interface{}, args *emul.SendArgs, wg *sync.WaitGroup) (err error) {
	var i uint
	var filePath string
	var filelist []string

	defer wg.Done()
	defer close(out)

	switch args.SendType {
	case emul.RepeatReq:
		// Prepare to send same repeat request.
		for i = 1; i <= args.Repeat; i++ {
			filePath = args.Data
			out <- emul.Request{SesID: SessionID, ReqID: uint64(i), FilePath: filePath}
			logger.Out(logrus.DebugLevel, logrus.Fields{"filePath": filePath}, "Enqueued request.")
		}

	case emul.MultiReq:
		// Use Glob to get file list by pattern.
		filelist, err = filepath.Glob(args.SrcDir + "/" + args.Wildcard)
		if err != nil {
			log.Error("Cannot get files by pattern", err)
			return
		}

		if args.Repeat == 0 {
			// If repeat flag is not set, send all enumerated files.
			args.Repeat = uint(len(filelist))
		}

		if len(filelist) == 0 {
			logger.Out(logrus.DebugLevel,
				logrus.Fields{"args.RequestDir": args.SrcDir, "args.Pattern": args.Wildcard},
				"Empty request directory, no requests to send.")
		} else {
			var j int
			for i = 1; i <= args.Repeat; {
				for _, filePath := range filelist {
					if j == len(filelist) && i == 0 {
						// No files, only directories.
						i = args.Repeat
						break
					} else if i > args.Repeat {
						// All requests sent.
						break
					}

					if stat, err := os.Stat(filePath); err == nil {
						if !stat.IsDir() {
							logger.Out(logrus.DebugLevel, logrus.Fields{"filePath": filePath}, "Enqueue request.")
							out <- emul.Request{SesID: SessionID, ReqID: uint64(i), FilePath: filePath}
							i++
							logger.Out(logrus.DebugLevel, nil, "Enqueued request.")
						} else {
							logger.Out(logrus.DebugLevel, logrus.Fields{"filePath": filePath}, "Skipping directory.")
						}
					} else {
						logger.Out(logrus.WarnLevel, logrus.Fields{"filePath": filePath, "error": err}, "Cannot stat file.")
					}
					j++
				}
			}
		}

	default:
		logger.Out(logrus.DebugLevel, logrus.Fields{"args.SendType": args.SendType}, "Invalide send type to prepare requests.")
	}

	return
}
