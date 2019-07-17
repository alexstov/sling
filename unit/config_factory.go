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

package unit

import (
	"github.com/4alexey/sling/conf"
	"github.com/4alexey/sling/util"
)

// NewConfig creates new configuraiton.
func NewConfig() conf.SlingConfig {
	return conf.SlingConfig{File: "", Dir: "fdsa", Wildcard: "*", SaveReqDir: "", SaveResDir: "", Repeat: 4}
}

// NewConsoleConfig crates new console configuration.
func NewConsoleConfig() conf.Console {
	return conf.Console{Level: 5, Flat: false, DisableColors: false, FullTimestamp: true}
}

// NewLogConfig crates new log configuration.
func NewLogConfig() conf.Log {
	return conf.Log{LogDir: "dir", LogFile: "file", Level: 5, LogReqDir: "logReq",
		LogResDir: "logRes", DisableColors: true, FullTimestamp: true}
}

// NewThrottleDefaultConfig crates new log configuration.
func NewThrottleDefaultConfig() conf.Throttle {
	return conf.Throttle{CxnNum: util.MaxUint, SleepMs: util.MaxUint, RateSec: util.MaxUint, RateMin: util.MaxUint, TmoSec: util.MaxUint, TmoRdS: util.MaxUint, TmoWrS: util.MaxUint}
}
