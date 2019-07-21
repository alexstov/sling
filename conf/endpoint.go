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

package conf

import (
	"fmt"
	"strings"
)

// ClientType enum
type ClientType int

const (
	// UnknownClient unknown endpoint type
	UnknownClient ClientType = iota
	// TCP type
	TCP
	// HTTPPost type
	HTTPPost
)

// Endpoint configuration
type Endpoint struct {
	Address string
	Port    uint
	Type    ClientType
}

// ParseClinetType parses string to ClinetType
func ParseClinetType(str string) ClientType {
	if strings.EqualFold(str, fmt.Sprintf("%s", TCP)) {
		return TCP
	} else if strings.EqualFold(str, fmt.Sprintf("%s", HTTPPost)) {
		return HTTPPost
	} else {
		return UnknownClient
	}
}
