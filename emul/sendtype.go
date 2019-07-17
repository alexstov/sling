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

// SendType request type.
type SendType int

const (
	// UnknownReq unknown send type.
	UnknownReq SendType = iota
	// SingleReq a single send request.
	SingleReq
	// RepeatReq repeat the same request.
	RepeatReq
	// MultiReq multiple requests.
	MultiReq
)

func (s SendType) String() string {
	return [...]string{"UnknownReq", "SingleReq", "RepeatReq", "MultiReq"}[s]
}
