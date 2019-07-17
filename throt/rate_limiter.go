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

package throt

import (
	"context"
	"sort"
	"time"

	"golang.org/x/time/rate"
)

// Limiter limits the rate per second and/or minute
type Limiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

// MultiLimitArgs used to configure the limiter.
type MultiLimitArgs struct {
	CxnNum  uint
	RateSec uint
	RateMin uint
}

// NewMultiLimiter limits the rate using multiple limiters.
func NewMultiLimiter(args *MultiLimitArgs) (limiter Limiter, err error) {
	secondLimit := rate.NewLimiter(Per(int(args.RateSec), time.Second), int(args.CxnNum))
	minuteLimit := rate.NewLimiter(Per(int(args.RateMin), time.Minute), int(args.CxnNum))
	limiters := []Limiter{secondLimit, minuteLimit}

	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	lim := MultiLimiter{limiters: limiters}
	lim.limiter = &lim
	return lim.limiter, nil
}

// MultiLimiter throttles transactions
type MultiLimiter struct {
	limiters []Limiter
	limiter  Limiter
}

// Wait blocks to limit the rate by using every encapsulted limiter.
func (l *MultiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}

	return nil
}

// Limit returns the most restrictive rate limiter.
func (l *MultiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

// Per sets rate limit. Returns rate.Inf if duration is zero.
func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount)) // TODO: bug
}
