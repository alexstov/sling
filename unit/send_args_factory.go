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

import "github.com/alexstov/sling/emul"

// NewSendArgs creates new configuraiton.
func NewSendArgs() emul.SendArgs {
	var sendArgs = emul.SendArgs{
		Data:       "string",
		SrcDir:     "string",
		Wildcard:   "string",
		SaveReqDir: "string",
		SaveResDir: "string",
		Repeat:     2,
		CxnNum:     2,
		SleepMs:    1,
		Address:    "string",
		Port:       432432,
		RateSec:    32,
		RateMin:    65,
		SendType:   emul.MultiReq,
	}

	return sendArgs
}
