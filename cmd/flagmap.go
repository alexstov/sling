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
	"fmt"
	"os"

	"github.com/4alexey/sling/conf"
	"github.com/4alexey/sling/emul"
	"github.com/4alexey/sling/sio"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// ID enum
type ID int

const (
	// CmdUnknown unknown command
	CmdUnknown ID = iota
	// CmdRoot command
	CmdRoot
	// CmdSend send command
	CmdSend
	// CmdView view command
	CmdView
)

var cmdUse = [...]string{
	"unknown",
	"root",
	"send",
	"view",
}

// Flags has all command flags.
type Flags struct {
	Command    *cobra.Command
	Map        map[FlagID]*Flag
	Explicit   map[FlagID]*Flag
	Filer      sio.Filer
	Flagmapper Flagmapper
}

// NewCmdFlags adds all command flags for the specific command by ID.
func NewCmdFlags(cmd *cobra.Command, id ID, slconf *conf.SlingConfig) (flagmapper Flagmapper, err error) {
	var filer sio.Filer
	if filer, err = sio.NewFiler(logger); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create flagmap filer.")
		return
	}

	flagmapper = NewFlagMap(cmd, filer)
	switch id {
	case CmdRoot:
		flagmapper.Add(NewFlagUint(LogLvl, sconf.Log.Level), true)
		flagmapper.Add(NewFlagUint(ConLvl, sconf.Console.Level), true)
		flagmapper.Add(NewFlagBool(ConFlat, sconf.Console.Flat), true)
	case CmdSend:
		flagmapper.Add(NewFlagStr(File, sconf.File), false)
		flagmapper.Add(NewFlagUint(Repeat, sconf.Repeat), false)
		flagmapper.Add(NewFlagStr(Dir, sconf.Dir), false)
		flagmapper.Add(NewFlagStr(Wildcard, sconf.Wildcard), false)
		flagmapper.Add(NewFlagUint(CxnNum, sconf.Throttle.CxnNum), false)
		flagmapper.Add(NewFlagUint(SleepMs, sconf.Throttle.SleepMs), false)
		flagmapper.Add(NewFlagStr(SaveReqDir, sconf.SaveReqDir), false)
		flagmapper.Add(NewFlagStr(SaveResDir, sconf.SaveResDir), false)
		flagmapper.Add(NewFlagUint(RateSec, sconf.Throttle.RateSec), false)
		flagmapper.Add(NewFlagUint(RateMin, sconf.Throttle.RateMin), false)
		flagmapper.Add(NewFlagBool(LogHis, sconf.Log.Histogram), false)
		flagmapper.Add(NewFlagBool(ConHis, sconf.Console.Histogram), false)
		flagmapper.Add(NewFlagUint(TmoSec, sconf.Throttle.TmoSec), false)
		flagmapper.Add(NewFlagUint(TmoRdS, sconf.Throttle.TmoRdS), false)
		flagmapper.Add(NewFlagUint(TmoWrS, sconf.Throttle.TmoWrS), false)
		flagmapper.Add(NewFlagBool(CxnLim, sconf.Throttle.CxnLim), false)
		flagmapper.Add(NewFlagBool(SaveReq, sconf.SaveReq), false)
		flagmapper.Add(NewFlagBool(SaveRes, sconf.SaveRes), false)
		flagmapper.Add(NewFlagUint(TmoCxn, sconf.Throttle.TmoCxn), false)
		flagmapper.Add(NewFlagUint(Endpoint, sconf.EndpointIndex), false)
		flagmapper.Add(NewFlagStr(Address, sconf.Endpoints[sconf.EndpointIndex].Address), false)
		flagmapper.Add(NewFlagUint(Port, sconf.Endpoints[sconf.EndpointIndex].Port), false)
		flagmapper.Add(NewFlagStr(CltType, fmt.Sprintf("%s", sconf.Endpoints[sconf.EndpointIndex].Type)), false)
	}

	flagmapper.SetExplicit()
	return flagmapper, nil
}

// NewFlagMap creates a new Flagmapper instance.
func NewFlagMap(cmd *cobra.Command, filer sio.Filer) Flagmapper {
	flagmap := &Flags{
		Command:  cmd,
		Map:      make(map[FlagID]*Flag),
		Explicit: make(map[FlagID]*Flag),
	}
	flagmap.Filer = filer
	flagmap.Flagmapper = flagmap
	return flagmap
}

// GetFlagmap returns the flag map.
func (fs Flags) GetFlagmap() map[FlagID]*Flag {
	return fs.Map
}

// GetExplicit returns the explicit flag map.
func (fs Flags) GetExplicit() map[FlagID]*Flag {
	return fs.Explicit
}

// SetExplicit finds the flags set explicitly and addes them to the Explicit map.
func (fs Flags) SetExplicit() {
	var explicit bool
	var err error
	for id, flag := range fs.Map {
		explicit, err = flag.Changed()
		if err != nil {
			logger.Out(logrus.ErrorLevel, logrus.Fields{"id": id, "flag": flag}, "Invalid flag.")
			continue
		}
		if explicit {
			fs.Explicit[flag.ID] = flag
			if flag.ID == ConFlat {
				fs.Map[ConFlat].FireEvent(Events[ConsoleSetFlat], fs.Map[ConFlat].Value.(*BoolVal).Value)
			} else if flag.ID == LogLvl {
				fs.Map[LogLvl].FireEvent(Events[LogSetLevel], logrus.Level(fs.Map[LogLvl].Value.(*UintVal).Value))
			} else if flag.ID == ConLvl {
				fs.Map[ConLvl].FireEvent(Events[ConsoleSetLevel], logrus.Level(fs.Map[ConLvl].Value.(*UintVal).Value))
			} else if flag.ID == Endpoint {
				fs.Map[Endpoint].FireEvent(Events[ActiveEndpointSet], fs.Map[Endpoint].Value.(*UintVal).Value)
			}

			logger.Out(logrus.DebugLevel, logrus.Fields{"id": id, "flag": flag}, "Flag set explicitly.")
		}
	}
}

// ResolveSendArgs resolves send arguments from command flags.
func (fs Flags) ResolveSendArgs(args *emul.SendArgs) (err error) {
	var fi os.FileInfo

	// Resolve file flag.
	if flag, ok := fs.Explicit[File]; ok {
		// File is set explicitly.
		args.SendType = emul.SingleReq

		// Get the file attributes for file or directory.
		if fi, err = os.Stat(flag.Value.(*StrVal).Value); os.IsNotExist(err) {
			// Only file name is passed. Build the full path.
			// NOTE: SrcDir may have been passed explicitly too.
			// TODO: validate file atrributes here?
			if fi, err = os.Stat(fs.Map[Dir].Value.(*StrVal).Value); err != nil || fi == nil || !fi.IsDir() {
				logger.Out(logrus.ErrorLevel, logrus.Fields{"id": flag.ID, "flag": fs.Map[Dir]}, "Invalid directory.")
				err = errors.Wrap(err, "os.Stat")
				return
			}

			var filePath string
			if filePath, err = fs.Filer.BuildFilePath(fs.Map[Dir].Value.(*StrVal).Value, flag.Value.(*StrVal).Value); err != nil {
				logger.Out(logrus.ErrorLevel, logrus.Fields{"id": flag.ID, "flag": fs.Map[Dir]}, "Invalid filepath.")
				err = errors.Wrap(err, "os.Stat")
				return
			}

			// Successfully built full file path using directory and filename.
			args.Data = filePath
			args.SendType = emul.SingleReq
		} else {
			switch mode := fi.Mode(); {
			case mode.IsRegular():
				args.Data = flag.Value.(*StrVal).Value
			}
		}
	} else if flag, ok := fs.Explicit[Dir]; ok {
		// SrcDir is set explicitly.
		args.SendType = emul.MultiReq

		if fi, err = os.Stat(flag.Value.(*StrVal).Value); err != nil || fi == nil || !fi.IsDir() {
			logger.Out(logrus.ErrorLevel, logrus.Fields{"id": flag.ID, "flag": fs.Map[Dir]}, "Invalid directory.")
			err = errors.Wrap(err, "os.Stat")
			return
		}
	}

	// Resolve repeat flag.
	if flag, ok := fs.Map[Repeat]; ok {
		rep := flag.Value.(*UintVal).Value

		if _, ok := fs.Explicit[Repeat]; ok {
			// Send only Repeat number of files for MultiReq.
			args.Repeat = uint(rep) // TODO: truncated?
		} else if args.SendType == emul.MultiReq {
			// Send all MultiReq files otherwise.
			args.Repeat = 0
		}

		if args.SendType == emul.SingleReq {
			args.Repeat = uint(rep) // TODO: truncated?
			if args.Repeat > 1 {
				// Change SingleReq to RepeatReq.
				args.SendType = emul.RepeatReq
			}
		}
	}

	// Apply Wildcard
	if flag, ok := fs.Map[Wildcard]; ok {
		// Apply wildcard.
		args.Wildcard = flag.Value.(*StrVal).Value
	}

	// Apply SrcDir
	if flag, ok := fs.Map[Dir]; ok {
		// Apply wildcard.
		args.SrcDir = flag.Value.(*StrVal).Value
	}

	if flag, ok := fs.Map[SleepMs]; ok {
		// Apply send delay.
		args.SleepMs = flag.Value.(*UintVal).Value
	}
	// Apply connectoin number and delay for RepeatReq and MultiReq
	if args.SendType == emul.RepeatReq || args.SendType == emul.MultiReq {
		if flag, ok := fs.Map[CxnNum]; ok {
			// Apply connection number.
			args.CxnNum = flag.Value.(*UintVal).Value
		}
	} else {
		args.CxnNum = 1
	}

	if flag, ok := fs.Map[Address]; ok {
		// Apply IP address.
		args.Address = flag.Value.(*StrVal).Value
	}

	if flag, ok := fs.Map[Port]; ok {
		// Apply port number.
		args.Port = flag.Value.(*UintVal).Value
	}

	// TODO: change to uint32
	if flag, ok := fs.Map[RateSec]; ok {
		// Apply port number.
		args.RateSec = flag.Value.(*UintVal).Value
	}

	// TODO: change to uint32
	if flag, ok := fs.Map[RateMin]; ok {
		// Apply port number.
		args.RateMin = flag.Value.(*UintVal).Value
	}

	if flag, ok := fs.Map[TmoSec]; ok {
		// Apply timeouts.
		args.TmoSec = flag.Value.(*UintVal).Value
	}
	if flag, ok := fs.Map[TmoRdS]; ok {
		args.TmoRdS = flag.Value.(*UintVal).Value

	}
	if flag, ok := fs.Map[TmoWrS]; ok {
		args.TmoWrS = flag.Value.(*UintVal).Value
	}
	if flag, ok := fs.Map[TmoCxn]; ok {
		args.TmoCxn = flag.Value.(*UintVal).Value
	}
	if flag, ok := fs.Map[CxnLim]; ok {
		args.CxnLim = flag.Value.(*BoolVal).Value
	}
	if flag, ok := fs.Map[CltType]; ok {
		args.CltType = conf.ParseClinetType(flag.Value.(*StrVal).Value)
	}
	if flag, ok := fs.Map[SaveReq]; ok {
		args.SaveReq = flag.Value.(*BoolVal).Value
	}
	if flag, ok := fs.Map[SaveReqDir]; ok {
		args.SaveReqDir = flag.Value.(*StrVal).Value + "/" + SessionID
	}
	if flag, ok := fs.Map[SaveRes]; ok {
		args.SaveRes = flag.Value.(*BoolVal).Value
	}
	if flag, ok := fs.Map[SaveResDir]; ok {
		args.SaveResDir = flag.Value.(*StrVal).Value + "/" + SessionID
	}
	// if flag, ok := fs.Map[EptActv]; ok {
	// args. = flag.Value.(*StrVal).Value + "/" + SessionID
	// }

	return
}

// Add a flag to the command flag set.
func (fs Flags) Add(flag *Flag, persist bool) {
	fs.Map[flag.ID] = flag
	if flag.Flagset != nil {
		if flag.Flagset != &fs {
			panic(errors.New("Flag can belong to a single flagset only"))
		} else {
			logger.Out(logrus.DebugLevel, logrus.Fields{"flag": flag, "fs": fs}, "Flag already added to the flagset, do nothing.")
		}
	}
	flag.Flagset = &fs

	var cmdFlagSet *pflag.FlagSet
	if persist {
		cmdFlagSet = fs.Command.PersistentFlags()
	} else {
		cmdFlagSet = fs.Command.Flags()
	}

	switch flag.Type {
	case StringType:
		if _, ok := flag.Value.(*StrVal); ok {
			cmdFlagSet.StringVarP(&flag.Value.(*StrVal).Value, flag.Name, flag.Shorthand, flag.Value.(*StrVal).Default, flag.Usage)
		} else {
			logger.Out(logrus.DebugLevel, logrus.Fields{"flag": flag, "fs": fs}, "Invalid flag value TODO.")
		}
	case UintType:
		if _, ok := flag.Value.(*UintVal); ok {
			cmdFlagSet.UintVarP(&flag.Value.(*UintVal).Value, flag.Name, flag.Shorthand, flag.Value.(*UintVal).Default, flag.Usage)
		} else {
			logger.Out(logrus.DebugLevel, logrus.Fields{"flag": flag, "fs": fs}, "Invalid flag value TODO.")
		}
	case BoolType:
		if _, ok := flag.Value.(*BoolVal); ok {
			cmdFlagSet.BoolVarP(&flag.Value.(*BoolVal).Value, flag.Name, flag.Shorthand, flag.Value.(*BoolVal).Default, flag.Usage)
		} else {
			logger.Out(logrus.DebugLevel, logrus.Fields{"flag": flag, "fs": fs}, "Invalid flag value TODO.")
		}
	}
}

// AddEvent adds flag event.
func (fs Flags) AddEvent(flagID FlagID, name interface{}, performing interface{}, parameters ...interface{}) (flag bool, err error) {
	return fs.Map[flagID].AddEvent(name, performing, parameters...)
}

// SetActiveEndpoint sets active endpoint
func (fs Flags) SetActiveEndpoint(eptIdx uint) {
	fs.Map[Endpoint].SetValue(eptIdx)
	fs.Map[Address].SetValue(sconf.Endpoints[eptIdx].Address)
	fs.Map[Port].SetValue(sconf.Endpoints[eptIdx].Port)
	fs.Map[CltType].SetValue(fmt.Sprintf("%s", sconf.Endpoints[eptIdx].Type))
}
