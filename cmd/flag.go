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

	"github.com/agoalofalife/event"
	"github.com/pkg/errors" //"errors"
)

// FlagID enum
type FlagID int

// FlagType enum
type FlagType int

const (
	// UnknownFlag unknown flag, " ", Unknown Flag.
	UnknownFlag FlagID = iota
	// Address endpoint IP, DNS name, or HTTP address, -a, --address
	Address
	// CltType network client type, TCP or HttpPost, -c, --cltType
	CltType
	// ConHis write histogram to console, -y, --conHis
	ConHis
	// CxnLim limit the number of concurrent connections, -l, cxnLim
	CxnLim
	// CxnNum number of concurrent connections, -n, --cxnNum
	CxnNum
	// Dir directory to send files from, -d, --dir
	Dir
	// Endpoint active endpoint index in SLINGCONFIG, zero-based, -i, --endpont
	Endpoint
	// File file path or filename to send, -f, --file
	File
	// LogHis write histogram to log file, -g, --logHis
	LogHis
	// Port endpoint port number, -p, --port
	Port
	// RateMin send rate per minute, -m, --rateMin
	RateMin
	// RateSec send rate per second, -s, --rateSec
	RateSec
	// Repeat send repeat count, -r, --repeat
	Repeat
	// SaveReq save requests, -q, --saveReq
	SaveReq
	// SaveReqDir directory to save requests, -k, --saveReqDir
	SaveReqDir
	// SaveRes save responses, -o, --saveRes
	SaveRes
	// SaveResDir directory to save response, -j, --saveResDir
	SaveResDir
	// SleepMs delay after each repeated request, -e, --sleepMs
	SleepMs
	// TmoCxn network client dial timeout, -u, --tmoCxn
	TmoCxn
	// TmoRdS network client timeout for Read calls, -v, tmoRdS
	TmoRdS
	// TmoSec network client timeout, -t, --tmoSec
	TmoSec
	// TmoWrS network client timeout for Write calls, -x, --tmoWrS
	TmoWrS
	// Wildcard file name matching wildcard, -w, --wildcard
	Wildcard
	// ConFlat set console flat output without timestamp and fields, - , --conFlat
	ConFlat
	// ConLvl console output level, --, conLvl
	ConLvl
	// LogLvl log output level, --, logLvl
	LogLvl
)

const (
	// UnknownType unknown flag type, " ", Unknown Type
	UnknownType FlagType = iota
	// StringType flag
	StringType
	// BoolType flag
	BoolType
	// UintType flag
	UintType
	// UintSliceType flag
	UintSliceType
)

var flagUsage = [...]string{
	"unknown flag is not used",
	"endpoint IP, DNS name, or HTTP address",
	"network client type, TCP or HttpPost",
	"write histogram to console",
	"limit the number of concurrent connections",
	"number of concurrent connections",
	"directory to send files from",
	"active endpoint index in SLINGCONFIG, zero-based",
	"file path or file name to send",
	"write histogram to log file",
	"endpoint port number",
	"send rate per minute",
	"send rate per second",
	"send repeat count",
	"save requests",
	"directory to save requests",
	"save responses",
	"directory to save response",
	"delay after each repeated request",
	"network client dial timeout",
	"read network client timeout",
	"network client timeout",
	"write network client timeout",
	"file name matching wildcard",
	"log output level",
	"console output level",
	"set console flat output without timestamp and fields",
}

// EventID enum
type EventID int

const (
	// UnknownEvent unknown event
	UnknownEvent EventID = iota
	// ConsoleSetFlat sets console output flat with no timestamp or fields
	ConsoleSetFlat
	// LogSetLevel sets log level
	LogSetLevel
	// ConsoleSetLevel sets console level
	ConsoleSetLevel
	// WildcardSet sets the wildcard flag
	WildcardSet
	// ActiveEndpointSet sent the active enpoint index
	ActiveEndpointSet
)

// Events event names
var Events = [...]string{
	"unknows_event",
	"console_set_flat",
	"log_set_level",
	"console_set_level",
	"wildcard_set",
	"active_enpoint_set",
}

// A Flag represents the state of a flag.
type Flag struct {
	ID        FlagID
	Name      string
	Shorthand string
	Usage     string
	Flagset   *Flags
	Type      FlagType
	Value     interface{}
	Event     *event.Dispatcher
}

// UintVal flag value
type UintVal struct {
	Value   uint
	Default uint
}

// BoolVal flag value
type BoolVal struct {
	Value   bool
	Default bool
}

// StrVal flag value
type StrVal struct {
	Value   string
	Default string
}

// NewFlag returns a new, empty flag set with the specified it, name, etc.
func NewFlag(id FlagID, flagType FlagType) *Flag {
	f := &Flag{
		ID:        id,
		Name:      fmt.Sprintf("%s", id),
		Shorthand: id.Flag(),
		Usage:     flagUsage[id],
		Flagset:   nil,
		Type:      flagType,
		Value:     nil,
		Event:     event.New(),
	}
	return f
}

// NewFlagBool returns a new bool flag.
func NewFlagBool(id FlagID, defaultValue bool) *Flag {
	flag := NewFlag(id, BoolType)
	flag.Value = &BoolVal{
		Default: defaultValue,
	}

	return flag
}

// NewFlagUint returns a new uint flag.
func NewFlagUint(id FlagID, defaultValue uint) *Flag {
	flag := NewFlag(id, UintType)
	flag.Value = &UintVal{
		Default: defaultValue,
	}

	return flag
}

// NewFlagStr returns a new string flag.
func NewFlagStr(id FlagID, defaultValue string) *Flag {
	flag := NewFlag(id, StringType)
	flag.Value = &StrVal{
		Default: defaultValue,
	}

	return flag
}

// Changed returns true if default value was changed.
func (f Flag) Changed() (bool, error) {
	if f.Flagset == nil {
		return false, errors.New("Flag does not have a valid flagset")
	}
	if f.Flagset.Command == nil {
		return false, errors.New("Flag does not have a valid command")
	}

	return f.Flagset.Command.Flags().Changed(f.Name), nil
}

// AddEvent adds flag event
func (f Flag) AddEvent(name interface{}, performing interface{}, parameters ...interface{}) (flag bool, err error) {
	return f.Event.Add(name, performing, parameters...)
}

// FireEvent fire flag event.
func (f Flag) FireEvent(event interface{}, parameters ...interface{}) (err error) {
	return f.Event.Fire(event, parameters...)
}

// EqualVal returns bool if value is equal.
func (f *Flag) EqualVal(val interface{}) bool {
	switch f.Type {
	case StringType:
		return f.Value.(*StrVal).Value == val.(string)
	case UintType:
		return f.Value.(*UintVal).Value == val.(uint)
	case BoolType:
		return f.Value.(*BoolVal).Value == val.(bool)
	}

	return false
}

// SetValue sets flag value.
func (f Flag) SetValue(val interface{}) {
	switch f.Type {
	case StringType:
		f.Value.(*StrVal).Value = val.(string)
	case UintType:
		f.Value.(*UintVal).Value = val.(uint)
	case BoolType:
		f.Value.(*BoolVal).Value = val.(bool)
	}

	return
}

// EqualDefault returns bool if value is equal.
func (f *Flag) EqualDefault(val interface{}) bool {
	switch f.Type {
	case StringType:
		return f.Value.(*StrVal).Default == val.(string)
	case UintType:
		return f.Value.(*UintVal).Default == val.(uint)
	case BoolType:
		return f.Value.(*BoolVal).Default == val.(bool)
	}

	return false
}
