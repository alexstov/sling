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

package cui_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/4alexey/sling/conf"
	"github.com/4alexey/sling/cui"
	"github.com/4alexey/sling/mock"
	"github.com/4alexey/sling/unit"
)

var _ = Describe("Console", func() {

	var (
		err         error
		t           testing.T
		mockCtrl    *gomock.Controller
		consoleConf conf.Console
		// ctx         context.Context
		mockLogger *mock.MockLogger
		consoler   cui.Consoler
	)

	BeforeEach(func() {
		// ctx = context.TODO()
		mockCtrl = gomock.NewController(&t)
		mockLogger = mock.NewMockLogger(mockCtrl)
		consoleConf = unit.NewConsoleConfig()

		mockLogger.EXPECT().Out(logrus.DebugLevel,
			logrus.Fields{"Level": consoleConf.Level,
				"Flat":          consoleConf.Flat,
				"DisableColors": consoleConf.DisableColors,
				"FullTimestamp": consoleConf.FullTimestamp},
			"Set console parameters from config.").Return(nil).Times(1)
		mockLogger.EXPECT().Out(logrus.DebugLevel,
			logrus.Fields{"flat": consoleConf.Flat},
			"Set console flat.").Return(nil).Times(1)

		consoler, err = cui.NewConsole(&consoleConf, mockLogger)

		Expect(err).Should(BeNil())
		Expect(consoler).ShouldNot(BeNil())
	})

	Describe("Consoler", func() {
		Context("SetFlat", func() {
			It("true, false", func() {
				defer GinkgoRecover()
				defer mockCtrl.Finish()

				for i := 0; i < 2; i++ {
					flat := true

					mockLogger.EXPECT().Out(logrus.DebugLevel,
						logrus.Fields{"flat": flat},
						"Set console flat.").Return(nil).Times(1)

					consoler.SetFlat(flat)

					flat = !flat
				}
			})
		})
	})
})
