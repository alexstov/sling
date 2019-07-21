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
	"context"
	"sync"

	metrics "github.com/rcrowley/go-metrics"
)

// Dispatcher sends requests to the endpoint.
type Dispatcher interface {
	SendReq(ctx context.Context, filePath string, args *SendArgs) (err error)
	MultiSend(ctx context.Context, req <-chan interface{}, args *SendArgs, wgSend *sync.WaitGroup) (err error)
	GetHisto() (histo metrics.Histogram, err error)
}
