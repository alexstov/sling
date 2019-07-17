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

package slog_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/4alexey/sling/conf"
	. "github.com/4alexey/sling/slog"
	"github.com/4alexey/sling/unit"
)

var _ = Describe("Log", func() {
	var (
		err      error
		t        testing.T
		mockCtrl *gomock.Controller
		logger   Logger
		// mockLogger *mock.MockLogger
		logConf conf.Log
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(&t)
		logConf = unit.NewLogConfig()
		// mockLogger = mock.NewMockLogger(mockCtrl)
	})

	Describe("Logger", func() {
		Context("new", func() {
			It("created successfully", func() {
				defer mockCtrl.Finish()

				logger, err = NewLogger()
				Expect(err).Should(BeNil())
				Expect(logger).ShouldNot(BeNil())
			})

			It("with valid logrus", func() {
				logger, err = NewLogger()
				logru := logger.GetLogger()
				Expect(logru).ShouldNot(BeNil())
				fmtr := logru.Formatter
				Expect(fmtr).ShouldNot(BeNil())
				txtfmtr := fmtr.(*logrus.TextFormatter)
				Expect(txtfmtr.DisableColors).To(Equal(true)) // TODO: read from mock config
				Expect(txtfmtr.FullTimestamp).To(Equal(true)) // TODO: read from mock config
			})
		})
	})

	Describe("Logger", func() {
		Context("configured", func() {
			It("from Sling.Log", func() {
				defer mockCtrl.Finish()

				logger, err = NewLogger()
				err = logger.Configure(&logConf)
				Expect(err).Should(BeNil())
				logrus := logger.GetLogger()
				Expect(logrus).ShouldNot(BeNil())
				Expect(uint(logrus.Level)).To(Equal(logConf.Level))
			})

			// It("with valid logrus", func() {
			// 	logger, err = NewLogger()
			// 	logru := logger.GetLogger()
			// 	Expect(logru).ShouldNot(BeNil())
			// 	fmtr := logru.Formatter
			// 	Expect(fmtr).ShouldNot(BeNil())
			// 	txtfmtr := fmtr.(*logrus.TextFormatter)
			// 	Expect(txtfmtr.DisableColors).To(Equal(true)) // TODO: read from mock config
			// 	Expect(txtfmtr.FullTimestamp).To(Equal(true)) // TODO: read from mock config
			// })
		})
	})
})
