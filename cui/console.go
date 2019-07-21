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

package cui

import (
	"fmt"
	"os"

	"github.com/alexstov/sling/conf"
	"github.com/alexstov/sling/slog"
	"github.com/sirupsen/logrus"
)

// Console implements Consoler interface
type Console struct {
	log      *logrus.Logger
	conOut   ConsoleOutFunc
	logger   slog.Logger
	consoler Consoler
	hist     bool
}

// Out Colsoler interface implementation
func (con *Console) Out(level logrus.Level, fields logrus.Fields, args ...interface{}) (err error) {
	return con.conOut(level, fields, args...)
}

// NewConsole creates new console.
func NewConsole(cconf *conf.Console, logger slog.Logger) (con Consoler, err error) {
	conlog := logrus.New()
	conlog.Out = os.Stdout
	conlog.SetLevel(logrus.Level(cconf.Level))
	conlog.SetFormatter(&slog.SlingTextFormatter{
		DisableColors:   cconf.DisableColors,
		TimestampFormat: cconf.TimestampFormat,
		FullTimestamp:   cconf.FullTimestamp,
		ForceFormatting: true,
	})

	theCon := Console{log: conlog}
	theCon.logger = logger
	theCon.SetFlat(cconf.Flat)
	theCon.consoler = &theCon
	theCon.logger.Out(logrus.DebugLevel,
		logrus.Fields{"Level": cconf.Level,
			"Flat":          cconf.Flat,
			"DisableColors": cconf.DisableColors,
			"FullTimestamp": cconf.FullTimestamp},
		"Set console parameters from config.")

	return theCon.consoler, nil
}

// GetLogger returns encapsulated logrus instance
func (con *Console) GetLogger() (logrus *logrus.Logger) {
	return con.log
}

// OutLogAndConsole write to log and console.
func (con *Console) OutLogAndConsole(level logrus.Level, fields logrus.Fields, args ...interface{}) (err error) {
	con.logger.Out(level, fields, args...)
	con.Out(level, fields, args...)

	return nil
}

// SetFlat sets the console flat.
func (con *Console) SetFlat(consoleFlat bool) {
	if consoleFlat {
		con.conOut = func(level logrus.Level, fields logrus.Fields, a ...interface{}) (err error) {
			if con.log.GetLevel() >= logrus.Level(level) {
				_, err = fmt.Print(fmt.Sprintln(a...))
			}
			return err
		}
	} else {
		con.conOut = func(level logrus.Level, fields logrus.Fields, a ...interface{}) (err error) {
			switch level {
			case logrus.PanicLevel:
				con.log.WithFields(fields).Panicln(a...)
			case logrus.FatalLevel:
				con.log.WithFields(fields).Fatalln(a...)
			case logrus.ErrorLevel:
				con.log.WithFields(fields).Errorln(a...)
			case logrus.WarnLevel:
				con.log.WithFields(fields).Warnln(a...)
			case logrus.InfoLevel:
				con.log.WithFields(fields).Infoln(a...)
			case logrus.DebugLevel:
				con.log.WithFields(fields).Debugln(a...)
			case logrus.TraceLevel:
				con.log.WithFields(fields).Traceln(a...)
			}
			return nil
		}
	}
	con.logger.Out(logrus.DebugLevel,
		logrus.Fields{"flat": consoleFlat},
		"Set console flat.")
}

// SetLevel sets the console flat.
func (con *Console) SetLevel(level logrus.Level) {
	con.log.SetLevel(level)
}
