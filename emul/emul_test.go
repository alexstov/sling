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

package emul_test

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/alexstov/sling/emul"
	"github.com/alexstov/sling/mock"
	"github.com/alexstov/sling/sio"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo" //"errors"

	// . "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("Emul", func() {
	var (
		ctx             context.Context
		t               testing.T
		mockCtrl        *gomock.Controller
		mockDispatcher  *mock.MockDispatcher
		mockClient      *mock.MockClient
		mockFiler       *mock.MockFiler
		mockLimiter     *mock.MockLimiter
		mockLogger      *mock.MockLogger
		mockHisto       *mock.MockHistogram
		mockConsoler    *mock.MockConsoler
		sendArgs        emul.SendArgs
		sesID           string
		testEmul        *emul.Emul
		testFileContent []byte
	)

	BeforeEach(func() {
		ctx = context.TODO()
		rand.Seed(time.Now().UnixNano())

		// Create all mocks.
		mockCtrl = gomock.NewController(&t)
		mockFiler = mock.NewMockFiler(mockCtrl)
		mockClient = mock.NewMockClient(mockCtrl)
		mockDispatcher = mock.NewMockDispatcher(mockCtrl)
		mockLimiter = mock.NewMockLimiter(mockCtrl)
		mockLogger = mock.NewMockLogger(mockCtrl)
		mockConsoler = mock.NewMockConsoler(mockCtrl)
		mockHisto = mock.NewMockHistogram(mockCtrl)

		// Create emul with default settings and mock interfaces.
		testEmul = &emul.Emul{Dispatcher: mockDispatcher, Client: mockClient, Filer: mockFiler, Limiter: mockLimiter, Logger: mockLogger, Histogram: mockHisto, Consoler: mockConsoler}

		// Deafult send arguments.
		sendArgs.SendType = emul.RepeatReq
		sendArgs.Repeat = 1
		sendArgs.SleepMs = 10
		sendArgs.CxnNum = 2
		sendArgs.SendType = emul.RepeatReq
		sendArgs.Repeat = 1
		sendArgs.Address = "127.0.0.1"
		sendArgs.Port = 9897
		sendArgs.ReqID = 1
		sendArgs.Data = "datafilepath.dat"

		t := time.Now()
		sesID = fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond())

	})

	Describe("SendReq", func() {
		Context("RepeatReq, single request", func() {
			It("normal flow.", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				// Set arguments.
				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1

				testFileContent = []byte("mock file content")
				mockFiler.EXPECT().DetermineContentType(sendArgs.Data).Return(sio.UnknownType, nil).Times(int(sendArgs.Repeat))
				mockFiler.EXPECT().ReadFile(sendArgs.Data).Return(testFileContent, nil).Times(int(sendArgs.Repeat))
				mockLimiter.EXPECT().Wait(ctx).Return(nil).Times(int(sendArgs.Repeat))
				if sendArgs.SaveReq {
					mockLogger.EXPECT().Out(logrus.InfoLevel, gomock.Any(), "Saved request to a file.")
				}
				mockClient.EXPECT().Write(testFileContent, gomock.Any()).Return(nil).Times(int(sendArgs.Repeat))
				mockLogger.EXPECT().Out(logrus.DebugLevel, nil, "Capturing Client execution stats.").Return(nil).Times(1)
				mockHisto.EXPECT().Update(gomock.Any()).Times(1)
				mockConsoler.EXPECT().OutLogAndConsole(logrus.InfoLevel, gomock.Any(), "Request sent successfully.").Return(nil).Times(1)

				testEmul.SetLogger(mockLogger)
				testEmul.SendReq(ctx, sendArgs.Data, &sendArgs)
			})

			It("normal flow, Gzip file.", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				// Set arguments.
				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1

				testFileContent = []byte("mock file content")
				mockFiler.EXPECT().DetermineContentType(sendArgs.Data).Return(sio.GzipType, nil).Times(int(sendArgs.Repeat))
				mockFiler.EXPECT().ReadArchive(sendArgs.Data).Return(testFileContent, nil).Times(int(sendArgs.Repeat))
				mockLimiter.EXPECT().Wait(ctx).Return(nil).Times(int(sendArgs.Repeat))
				if sendArgs.SaveReq {
					mockLogger.EXPECT().Out(logrus.InfoLevel, gomock.Any(), "Saved request to a file.")
				}
				mockClient.EXPECT().Write(testFileContent, gomock.Any()).Return(nil).Times(int(sendArgs.Repeat))
				mockLogger.EXPECT().Out(logrus.DebugLevel, nil, "Capturing Client execution stats.").Return(nil).Times(1)
				mockHisto.EXPECT().Update(gomock.Any()).Times(1)
				mockConsoler.EXPECT().OutLogAndConsole(logrus.InfoLevel, gomock.Any(), "Request sent successfully.").Return(nil).Times(1)

				testEmul.SetLogger(mockLogger)
				testEmul.SendReq(ctx, sendArgs.Data, &sendArgs)
			})

			It("normal flow, Zip file.", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				// Set arguments.
				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1

				testFileContent = []byte("mock file content")
				mockFiler.EXPECT().DetermineContentType(sendArgs.Data).Return(sio.GzipType, nil).Times(int(sendArgs.Repeat))
				mockFiler.EXPECT().ReadArchive(sendArgs.Data).Return(testFileContent, nil).Times(int(sendArgs.Repeat))
				mockLimiter.EXPECT().Wait(ctx).Return(nil).Times(int(sendArgs.Repeat))
				if sendArgs.SaveReq {
					mockLogger.EXPECT().Out(logrus.InfoLevel, gomock.Any(), "Saved request to a file.")
				}
				mockClient.EXPECT().Write(testFileContent, gomock.Any()).Return(nil).Times(int(sendArgs.Repeat))
				mockLogger.EXPECT().Out(logrus.DebugLevel, nil, "Capturing Client execution stats.").Return(nil).Times(1)
				mockHisto.EXPECT().Update(gomock.Any()).Times(1)
				mockConsoler.EXPECT().OutLogAndConsole(logrus.InfoLevel, gomock.Any(), "Request sent successfully.").Return(nil).Times(1)

				testEmul.SetLogger(mockLogger)
				testEmul.SendReq(ctx, sendArgs.Data, &sendArgs)
			})

			It("normal flow, save request.", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				// Set arguments.
				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1

				testFileContent = []byte("mock file content")
				mockFiler.EXPECT().DetermineContentType(sendArgs.Data).Return(sio.UnknownType, nil).Times(int(sendArgs.Repeat))
				mockFiler.EXPECT().ReadFile(sendArgs.Data).Return(testFileContent, nil).Times(int(sendArgs.Repeat))
				mockLimiter.EXPECT().Wait(ctx).Return(nil).Times(int(sendArgs.Repeat))
				if sendArgs.SaveReq {
					// Save request calls.
					saveReqFilepath := "/tmp/saveRequest.file"
					mockFiler.EXPECT().BuildFilePath(gomock.Any(), gomock.Any()).Return(saveReqFilepath, nil).Times(int(sendArgs.Repeat))
					mockFiler.EXPECT().CreateFile(saveReqFilepath).Return(gomock.Any(), nil).Times(int(sendArgs.Repeat))
					mockFiler.EXPECT().WriteFile(gomock.Any(), nil).Return(gomock.Any(), nil).Times(int(sendArgs.Repeat))
					mockLogger.EXPECT().Out(logrus.InfoLevel, gomock.Any(), "Saved request to a file.").Times(int(sendArgs.Repeat))
					mockFiler.EXPECT().CloseFile(gomock.Any()).Return(nil).Times(int(sendArgs.Repeat))
				}
				mockClient.EXPECT().Write(testFileContent, gomock.Any()).Return(nil).Times(int(sendArgs.Repeat))
				mockLogger.EXPECT().Out(logrus.DebugLevel, nil, "Capturing Client execution stats.").Return(nil).Times(1)
				mockHisto.EXPECT().Update(gomock.Any()).Times(1)
				mockConsoler.EXPECT().OutLogAndConsole(logrus.InfoLevel, gomock.Any(), "Request sent successfully.").Return(nil).Times(1)

				testEmul.SetLogger(mockLogger)
				testEmul.SendReq(ctx, sendArgs.Data, &sendArgs)
			})

			It("Clinet time out.", func() {
				// Defer asserts.
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				testEmul := &emul.Emul{Dispatcher: mockDispatcher, Client: mockClient, Filer: mockFiler, Limiter: mockLimiter, Logger: mockLogger, Histogram: mockHisto, Consoler: mockConsoler}
				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1
				sendArgs.SleepMs = 1000
				sendArgs.CxnNum = 2
				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1
				sendArgs.Address = "127.0.0.1"
				sendArgs.Port = 9897
				filepath := "myfile.dat"
				sendArgs.TmoSec = 1
				ret := []byte("mock file content")
				mockFiler.EXPECT().DetermineContentType(filepath).Return(sio.UnknownType, nil).Times(int(sendArgs.Repeat))
				mockFiler.EXPECT().ReadFile(filepath).Return(ret, nil).Times(int(sendArgs.Repeat))
				mockClient.EXPECT().Write(ret, gomock.Any()).Return(nil).Times(int(sendArgs.Repeat))
				mockLimiter.EXPECT().Wait(ctx).Return(nil).Times(int(sendArgs.Repeat))
				mockLogger.EXPECT().Out(logrus.DebugLevel, nil, "Capturing Client execution stats.").Return(nil).Times(1)
				mockHisto.EXPECT().Update(gomock.Any()).Times(1)
				mockConsoler.EXPECT().OutLogAndConsole(logrus.InfoLevel, gomock.Any(), "Request sent successfully.").Return(nil).Times(1)

				testEmul.SetLogger(mockLogger)
				testEmul.SendReq(ctx, filepath, &sendArgs)
			})
		})
	})

	Describe("Dispatcher MultiSend", func() {
		Context("RepeatReq, single requst", func() {
			It("with invalid filepath.", func() {
				// Defer asserts
				defer mockCtrl.Finish()
				defer GinkgoRecover()

				sendArgs.SendType = emul.RepeatReq
				sendArgs.Repeat = 1
				sendArgs.SleepMs = 1000
				sendArgs.CxnNum = 2

				in := make(chan interface{}, sendArgs.Repeat)
				for i := uint(0); i < sendArgs.Repeat; i++ {
					in <- emul.Request{FilePath: "myfile.dat", SesID: sesID, ReqID: uint64(i)}
				}
				testEmul := &emul.Emul{Dispatcher: mockDispatcher, Client: mockClient, Filer: mockFiler, Consoler: mockConsoler}
				testEmul.SetLogger(mockLogger)

				var wg sync.WaitGroup
				wg.Add(1)
				testEmul.MultiSend(ctx, in, &sendArgs, &wg)
				wg.Wait()
			})
		})
	})
})
