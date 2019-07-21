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

package slog

import (
	"os"

	"github.com/alexstov/sling/conf"
	"github.com/sirupsen/logrus"
	//"errors"
)

// The log level befor config is applied.
const startupLogLevel = logrus.WarnLevel

// Log implements Logger interface
type Log struct {
	log    *logrus.Logger
	logger Logger
}

// Out Logger interface implementation
func (log *Log) Out(level logrus.Level, fields logrus.Fields, args ...interface{}) (err error) {
	switch level {
	case logrus.PanicLevel:
		log.log.WithFields(fields).Panicln(args...)
	case logrus.FatalLevel:
		log.log.WithFields(fields).Fatalln(args...)
	case logrus.ErrorLevel:
		log.log.WithFields(fields).Errorln(args...)
	case logrus.WarnLevel:
		log.log.WithFields(fields).Warnln(args...)
	case logrus.InfoLevel:
		log.log.WithFields(fields).Infoln(args...)
	case logrus.DebugLevel:
		log.log.WithFields(fields).Debugln(args...)
	case logrus.TraceLevel:
		log.log.WithFields(fields).Traceln(args...)
	}
	return
}

// NewLogger creates new logger
func NewLogger() (logger Logger, err error) {
	logru := logrus.New()
	logru.SetLevel(startupLogLevel)
	logru.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	theLogger := Log{log: logru}
	theLogger.logger = &theLogger
	return &theLogger, nil
}

// Configure configures logger
func (log *Log) Configure(logconf *conf.Log) (err error) {
	logger := log.GetLogger()

	logger.SetLevel(logrus.Level(logconf.Level))

	// os.O_WRONLY tells the computer you are only going to writo to the file, not read
	// os.O_CREATE tells the computer to create the file if it doesn't exist
	// os.O_APPEND tells the computer to append to the end of the file instead of overwritting or truncating it
	file, err := os.OpenFile(logconf.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.WithFields(logrus.Fields{"conf.Log.LogFile": logconf.LogFile, "error": err}).Errorln("Failed to open the log file.")
		return
	}
	logger.SetFormatter(&SlingTextFormatter{
		DisableColors:   logconf.DisableColors,
		TimestampFormat: logconf.TimestampFormat,
		FullTimestamp:   logconf.FullTimestamp,
		ForceFormatting: true,
	})
	logger.SetOutput(file)
	logger.WithFields(logrus.Fields{"conf.Log.LogFile": logconf.LogFile}).Debugln("Logger started.")

	return
}

// GetLogger returns encapsulated logrus instance
func (log *Log) GetLogger() (logrus *logrus.Logger) {
	return log.log
}

// SetLevel sets the console flat.
func (log *Log) SetLevel(level logrus.Level) {
	log.log.SetLevel(level)
}
