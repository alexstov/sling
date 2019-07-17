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

package emul

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/4alexey/sling/conf"
	"github.com/4alexey/sling/cui"
	"github.com/4alexey/sling/net"
	"github.com/4alexey/sling/sio"
	"github.com/4alexey/sling/slog"
	"github.com/4alexey/sling/throt"
	"github.com/pkg/errors" //"errors"
	"github.com/rcrowley/go-metrics"
	"github.com/sirupsen/logrus"
)

// Emul - Emulatator interface implementation.
type Emul struct {
	Dispatcher Dispatcher
	Client     net.Client
	Filer      sio.Filer
	Consoler   cui.Consoler
	Logger     slog.Logger
	Limiter    throt.Limiter
	Histogram  metrics.Histogram
}

// SendArgs send command arguments.
type SendArgs struct {
	Data            string
	SrcDir          string
	Wildcard        string
	SaveReqDir      string
	SaveReqFilepath string
	SaveResDir      string
	SaveResFilepath string
	Repeat          uint
	CxnNum          uint
	SleepMs         uint
	Address         string
	Port            uint
	RateSec         uint
	RateMin         uint
	SendType        SendType
	TmoSec          uint
	TmoRdS          uint
	TmoWrS          uint
	TmoCxn          uint
	CxnLim          bool
	CltType         conf.ClientType
	SaveReq         bool
	SaveRes         bool
	ReqID           uint64
}

// NewEmul creates new emul instance.
func NewEmul(clt net.Client, flr sio.Filer, con cui.Consoler, limiter throt.Limiter, logger slog.Logger, histo metrics.Histogram) (em *Emul, err error) {
	em = &Emul{Client: clt, Filer: flr, Consoler: con, Limiter: limiter, Logger: logger}
	em.Dispatcher = em
	em.Histogram = histo
	return em, nil
}

// MultiSend dispatches requests, controlling the number of concurrent connections.
func (em *Emul) MultiSend(ctx context.Context, in <-chan interface{}, args *SendArgs, wgSend *sync.WaitGroup) (err error) {
	defer wgSend.Done()

	var wg sync.WaitGroup
	res := make(chan interface{}, args.Repeat)

	if args.CxnLim {
		wg.Add(int(args.CxnNum))
		for w := 0; w < int(args.CxnNum); w++ {
			go em.dispatch(ctx, args, in, res, &wg)
		}
	} else {
		wg.Add(int(args.Repeat))
		for w := 0; w < int(args.Repeat); w++ {
			go em.dispatch(ctx, args, in, res, &wg)
		}
	}
	wg.Wait()
	return
}

// dispatch displatches input requests routed via inbound channel and sends responses
// through outbound channel. Send rate is adjusted using SleepMs.
func (em *Emul) dispatch(ctx context.Context, args *SendArgs, in <-chan interface{}, res chan<- interface{}, wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	for r := range in {
		args.ReqID = r.(Request).ReqID
		if err = em.Dispatcher.SendReq(ctx, r.(Request).FilePath, args); err != nil {
			// Log an error. Do not return, attempt to send all requests.
			err = errors.Wrap(err, "SendReq")
			em.Logger.Out(logrus.ErrorLevel, logrus.Fields{"filePath": "filePath", "error": err}, "Failed to send the request.")
		}
		time.Sleep(time.Duration(args.SleepMs) * time.Millisecond)
	}
	return
}

// SendReq sends a single request to destination.
func (em *Emul) SendReq(ctx context.Context, filePath string, args *SendArgs) (err error) {
	var contentType sio.ContentType
	var buf []byte

	// Determine request content type.
	contentType, err = em.Filer.DetermineContentType(filePath)
	if err != nil {
		em.Logger.Out(logrus.ErrorLevel, nil, "Unknown file content type", err)
		return
	}

	// Read request.
	switch contentType {
	case sio.GzipType:
		buf, err = em.Filer.ReadArchive(filePath)
		if err != nil {
			err = errors.Wrap(err, "readArchive(filepath)")
			return
		}
	case sio.ZipType:
		buf, err = em.Filer.ReadArchive(filePath)
		if err != nil {
			err = errors.Wrap(err, "readArchive(filepath)")
			return
		}
	case sio.UnknownType:
		buf, err = em.Filer.ReadFile(filePath)
		if err != nil {
			err = errors.Wrap(err, "io.ReadFile(filepath)")
			return
		}
	}

	// Limit the rate.
	if err = em.Limiter.Wait(ctx); err != nil {
		err = errors.Wrap(err, "em.Limiter.Wait(ctx)")
		return
	}

	// Send request and capture execution time.
	writeArgs := net.WriteArgs{IPAddress: args.Address,
		Port:            args.Port,
		TmoSec:          args.TmoSec,
		TmoRdS:          args.TmoRdS,
		TmoWrS:          args.TmoWrS,
		TmoCxn:          args.TmoCxn,
		ReqID:           args.ReqID,
		RequestFilepath: filePath,
		SaveReq:         args.SaveReq,
		SaveReqDir:      args.SaveReqDir,
		SaveRes:         args.SaveRes,
		SaveResDir:      args.SaveResDir,
		CltType:         args.CltType}

	// Save response callback, defined here to reuse by any client after receiving response.
	if args.SaveRes {
		writeArgs.SaveResCallback = func(filepath string, buff *bytes.Buffer) (err error) {
			var wrLen int
			var f *os.File
			f, err = em.Filer.CreateFile(filepath)
			if err != nil {
				err = errors.Wrap(err, "os.Create(filepath)")
				return
			}
			defer em.Filer.CloseFile(f)

			if wrLen, err = em.Filer.WriteFile(f, buff.Bytes()); err != nil {
				err = errors.Wrap(err, "f.Write(buff)")
				return
			}
			em.Logger.Out(logrus.InfoLevel, logrus.Fields{"filepath": filepath, "wrLen": wrLen}, "Saved request to a file.")
			return
		}
	}

	// Save the request.
	if args.SaveReq {
		if args.SaveReqFilepath, err = em.Filer.BuildFilePath(args.SaveReqDir, fmt.Sprintf("%03d", args.ReqID)+"."+filepath.Base(filePath)+".req"); err != nil {
			em.Logger.Out(logrus.ErrorLevel, logrus.Fields{"filepath": args.SaveReqFilepath}, "Cannot save the request.")
			err = errors.Wrap(err, "os.Stat")
		}

		go func(filepath string, buff []byte) (err error) {
			var wrLen int
			var f *os.File
			f, err = em.Filer.CreateFile(filepath)
			if err != nil {
				err = errors.Wrap(err, "os.Create(filepath)")
				return
			}
			defer em.Filer.CloseFile(f)

			if wrLen, err = em.Filer.WriteFile(f, buff); err != nil {
				err = errors.Wrap(err, "f.Write(buff)")
				return
			}
			em.Logger.Out(logrus.InfoLevel, logrus.Fields{"filepath": filepath, "wrLen": wrLen}, "Saved request to a file.")
			return
		}(args.SaveReqFilepath, buf)
	}

	// Send the request.
	start := time.Now()
	err = em.Client.Write(buf, &writeArgs)
	elapsed := int64(time.Since(start)) / int64(time.Millisecond)
	if err != nil {
		em.Consoler.OutLogAndConsole(logrus.ErrorLevel, logrus.Fields{"FilePath": filePath, "error": err}, "Failed to send the request.")
	} else {
		em.Consoler.OutLogAndConsole(logrus.InfoLevel, logrus.Fields{"FilePath": filePath, "Length": len(buf)}, "Request sent successfully.")
	}

	// Update stats.
	if histo, errH := em.GetHisto(); errH != nil {
		em.Logger.Out(logrus.WarnLevel, logrus.Fields{"error": errH}, "Cannot capture Client execution stats.")
	} else {
		em.Logger.Out(logrus.DebugLevel, nil, "Capturing Client execution stats.")
		histo.Update(elapsed)
	}

	return
}

// GetHisto implements interface method to returns histogram.
func (em *Emul) GetHisto() (histo metrics.Histogram, err error) {
	if em.Histogram == nil {
		return nil, errors.New("histogram is not created")
	}

	return em.Histogram, nil
}

// SetLogger sets Emul logger.
func (em *Emul) SetLogger(logger slog.Logger) {
	em.Logger = logger
}
