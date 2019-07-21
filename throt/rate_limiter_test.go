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

package throt_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sheerun/queue"

	. "github.com/alexstov/sling/throt"
)

var _ = Describe("RateLimiter", func() {

	var (
		err      error
		t        testing.T
		mockCtrl *gomock.Controller

		//slingConfig *conf.SlingConfig
		limArgs MultiLimitArgs
		limiter Limiter
		ctx     context.Context
		q       *queue.Queue
	)

	BeforeEach(func() {
		// ctx = context.TODO()
		mockCtrl = gomock.NewController(&t)
		limArgs = MultiLimitArgs{CxnNum: 3, RateSec: 100, RateMin: 6000}
		limiter, err = NewMultiLimiter(&limArgs)
		Expect(err).Should(BeNil())
		Expect(limiter).ShouldNot(BeNil())
		ctx = context.Background()
		q = queue.New()

	})

	Describe("Limiter", func() {
		Context("Wait", func() {
			It("limit transaction rate per second", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				for i := 0; i < 600; i++ {
					q.Append(i)
				}
				q.Append(nil)

				// Assert initial queue length.
				Expect(601).To(Equal(q.Length()))

				timer := time.NewTimer(time.Duration(5 * time.Second))
				go func() {
					for i := 0; i < 600; i++ {
						limiter.Wait(ctx)
						q.Pop()
					}
				}()
				<-timer.C

				// Assert remaining message number is in expected range.
				Expect(q.Length()).Should(SatisfyAll(
					BeNumerically(">=", 101-3),
					BeNumerically("<=", 101)))
			})

			It("limit transaction rate per minute", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				// Initial burst is 100, rate is 200 / 60 = 3 per sec.
				limArgs = MultiLimitArgs{CxnNum: 100, RateSec: 400, RateMin: 200}
				limiter, err = NewMultiLimiter(&limArgs)
				Expect(err).Should(BeNil())
				Expect(limiter).ShouldNot(BeNil())

				for i := 0; i < 600; i++ {
					q.Append(i)
				}
				q.Append(nil)

				// Assert initial queue length.
				Expect(601).To(Equal(q.Length()))

				// Expect queue depth 601 - 100 (init. burst) - 5 * 3 (per.seq)
				timer := time.NewTimer(time.Duration(5 * time.Second))
				go func() {
					for i := 0; i < 600; i++ {
						limiter.Wait(ctx)
						q.Pop()
					}
				}()
				<-timer.C

				// Assert remaining message number is in expected range.
				Expect(q.Length()).Should(SatisfyAll(
					BeNumerically(">=", 601-100-5*3-3),
					BeNumerically("<=", 601-100-5*3)))
			})
		})
	})
})
