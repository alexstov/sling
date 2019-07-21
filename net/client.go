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

package net

import "github.com/alexstov/sling/conf"

// WriteArgs write method arguments.
type WriteArgs struct {
	IPAddress       string
	Port            uint
	TmoSec          uint
	TmoRdS          uint
	TmoWrS          uint
	TmoCxn          uint
	RequestFilepath string
	SaveReq         bool
	SaveReqDir      string
	SaveReqFilepath string
	SaveRes         bool
	SaveResDir      string
	SaveResFilepath string
	SesID           string
	ReqID           uint64
	CltType         conf.ClientType
	SaveResCallback SaveToFileFunc
}

// Client sends requests to the endpoint.
type Client interface {
	// Write(ipaddress string, port uint, msg []byte) (err error)
	Write(msg []byte, args *WriteArgs) (err error)
}
