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

package cmd_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/4alexey/sling/cmd"
)

var defValMap map[FlagID]interface{}
var actualFlags map[FlagID]*Flag
var expectedFlags map[FlagID]*Flag

var _ = Describe("Flag", func() {
	var (
	// defaultRepeat       = uint(3)
	// repeatFlag          *Flag
	// defaultLogLvl       = uint(4)
	// logLvlFlag          *Flag
	// defaultConLvl       = uint(3)
	// conLvlFlag          *Flag
	// defaultConFlat      = true
	// conFlatFlag         *Flag
	// defaultCxnNum       = uint(9)
	// cxnNumFlag          *Flag
	// defaultSleepMs      = uint(10)
	// sleepMsFlag         *Flag
	// defaultSaveReqDir   = "defaultSaveReqDir"
	// saveReqDirFlag      *Flag
	// defaultSaveResDir   = "defaultSaveResDir"
	// saveResDirFlag      *Flag
	// defaultIP           = "193.45.435.23"
	// ipFlag              *Flag
	// defaultPort         = uint(8086)
	// portFlag            *Flag
	// defaultRateSec      = uint(435)
	// rateSecFlag         *Flag
	// defaultRateMin      = uint(4553)
	// rateMinFlag         *Flag
	// defaultLogHis       = true
	// logHisFlag          *Flag
	// defaultConHis       = true
	// conHisFlag          *Flag
	// defaultTimeoutSec   = uint(30)
	// timeoutRdSecFlag    *Flag
	// defaultRdTimeoutSec = uint(10)
	// timeoutWrSecFlag    *Flag
	// defaultWrTimeoutSec = uint(25)
	// timeoutSecFlag      *Flag
	// defaultCxnLim       = true
	// cxnLimFlag          *Flag
	// defaultCltType      = "TCP"
	// cltTypeFlag         *Flag
	// defaultSaveReq      = true
	// saveReqFlag         *Flag
	// defaultSaveRes      = true
	// saveResFlag         *Flag
	// dafaultTmoCxn       = uint(10)
	// tmoCxnFlag          *Flag
	)

	BeforeEach(func() {
		defValMap = map[FlagID]interface{}{
			File:       "defaultFile.txt",
			Dir:        "/tmp/Default",
			Wildcard:   "*.txt",
			Repeat:     "2",    //uint(2),
			LogLvl:     "3",    //uint(3),
			ConLvl:     "4",    //uint(4),
			ConFlat:    "true", //true,
			CxnNum:     "5",    //uint(5),
			SleepMs:    "30",   //uint(30),
			SaveReqDir: "/tmp/SaveReqDir",
			SaveResDir: "/tmp/SaveResDir/",
			Address:    "127.0.0.1",
			Port:       "21345", //uint(21345),
			RateSec:    "100",   //uint(100),
			RateMin:    "6000",  //uint(6000),
			LogHis:     "true",  //true,
			ConHis:     "true",  //true,
			TmoSec:     "11",    //uint(10),
			TmoRdS:     "21`",   //uint(20),
			TmoWrS:     "22",    //uint(20),
			CxnLim:     "true",  //true,
			CltType:    "TCP",
			SaveReq:    "true", //true,
			SaveRes:    "true", //true,
			TmoCxn:     "13",   //uint(10),
		}
		actualFlags = createFlags()
		expectedFlags = createExpectedFlags()
	})

	Describe("File flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[File], expectedFlags[File])
			})
		})
	})

	Describe("Dir flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[Dir], expectedFlags[Dir])
			})
		})
	})

	Describe("Wildcard flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[Wildcard], expectedFlags[Wildcard])
			})
		})
	})

	Describe("Repeat flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[Repeat], expectedFlags[Repeat])
			})
		})
	})

	Describe("LogLvl flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[LogLvl], expectedFlags[LogLvl])
			})
		})
	})

	Describe("ConLvl flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[ConLvl], expectedFlags[ConLvl])
			})
		})
	})

	Describe("ConFlat flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[ConFlat], expectedFlags[ConFlat])
			})
		})
	})

	Describe("CxnNum flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[CxnNum], expectedFlags[CxnNum])
			})
		})
	})

	Describe("SleepMs flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[SleepMs], expectedFlags[SleepMs])
			})
		})
	})

	Describe("SaveReqDir flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[SaveReqDir], expectedFlags[SaveReqDir])
			})
		})
	})

	Describe("SaveResDir flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[SaveResDir], expectedFlags[SaveResDir])
			})
		})
	})

	Describe("IP flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[Address], expectedFlags[Address])
			})
		})
	})

	Describe("Port flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[Port], expectedFlags[Port])
			})
		})
	})

	Describe("RateSec flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[RateSec], expectedFlags[RateSec])
			})
		})
	})

	Describe("RateMin flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[RateMin], expectedFlags[RateMin])
			})
		})
	})

	Describe("LogHis flag", func() {
		Context("fields", func() {
			It("write histogram to log file.", func() {
				assertFlagFields(actualFlags[LogHis], expectedFlags[LogHis])
			})
		})
	})

	Describe("ConHis flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[ConHis], expectedFlags[ConHis])
			})
		})
	})

	Describe("TmoSec flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[TmoSec], expectedFlags[TmoSec])
			})
		})
	})

	Describe("TmoReqS flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[TmoRdS], expectedFlags[TmoRdS])
			})
		})
	})

	Describe("TmoWrS flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[TmoWrS], expectedFlags[TmoWrS])
			})
		})
	})

	Describe("CxnLim flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[CxnLim], expectedFlags[CxnLim])
			})
		})
	})

	Describe("CltType flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[CltType], expectedFlags[CltType])
			})
		})
	})

	Describe("SaveReq flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[SaveReq], expectedFlags[SaveReq])
			})
		})
	})

	Describe("SaveRes flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[SaveRes], expectedFlags[SaveRes])
			})
		})
	})

	Describe("TmoCxn flag", func() {
		Context("fields", func() {
			It("set to default values", func() {
				assertFlagFields(actualFlags[TmoCxn], expectedFlags[TmoCxn])
			})
		})
	})
})

func assertFlagFields(flag *Flag, expected *Flag) {
	defer GinkgoRecover()

	Expect(flag.ID).To(Equal(expected.ID))
	Expect(flag.Name).To(Equal(expected.Name))
	Expect(flag.Shorthand).To(Equal(expected.Shorthand))
	Expect(flag.EqualVal("")).To(BeTrue())
	Expect(flag.Usage).To(Equal(expected.Usage))
	Expect(flag.EqualDefault(defValMap[expected.ID])).To(BeTrue())
	Expect(flag.Flagset).To(BeNil())
}

func createFlags() map[FlagID]*Flag {
	flagMap := make(map[FlagID]*Flag)

	// File used for filenmane, "f", file.
	flagMap[File] = NewFlagStr(File, fmt.Sprint(defValMap[File]))
	// Dir used for directory, "d", dir.
	flagMap[Dir] = NewFlagStr(Dir, fmt.Sprint(defValMap[Dir]))
	// Wildcard used to pass wildcard for Dir, "w", wildcard.
	flagMap[Wildcard] = NewFlagStr(Wildcard, fmt.Sprint(defValMap[Wildcard]))
	// Repeat passes number of requests to send, "r", repeat.
	flagMap[Repeat] = NewFlagStr(Repeat, fmt.Sprint(defValMap[Repeat]))
	// LogLvl passes log level, "l", logLvl.
	flagMap[LogLvl] = NewFlagStr(LogLvl, fmt.Sprint(defValMap[LogLvl]))
	// ConLvl passes log level for console output, "c", conLvl.
	flagMap[ConLvl] = NewFlagStr(ConLvl, fmt.Sprint(defValMap[ConLvl]))
	// ConFlat used for flat console output without timestamp and fields., "t", conFlat.
	flagMap[ConFlat] = NewFlagStr(ConFlat, fmt.Sprint(defValMap[ConFlat]))
	// CxnNum number of connections for repeat or multisend, "x", cxnNum.
	flagMap[CxnNum] = NewFlagStr(CxnNum, fmt.Sprint(defValMap[CxnNum]))
	// SleepMs delay milliseconds for repeat or multisend, "e", sleepMs.
	flagMap[SleepMs] = NewFlagStr(SleepMs, fmt.Sprint(defValMap[SleepMs]))
	// SaveReqDir directory to save itermediary requests, "q", saveReqDir.
	flagMap[SaveReqDir] = NewFlagStr(SaveReqDir, fmt.Sprint(defValMap[SaveReqDir]))
	// SaveResDir directory to save responses, "v", saveResDir.
	flagMap[SaveResDir] = NewFlagStr(SaveResDir, fmt.Sprint(defValMap[SaveResDir]))
	// IP of the endpoint, "a", ip.
	flagMap[Address] = NewFlagStr(Address, fmt.Sprint(defValMap[Address]))
	// Port of the endpoint, "p", "port".
	flagMap[Port] = NewFlagStr(Port, fmt.Sprint(defValMap[Port]))
	// RateSec limits send per second rate "s", "rateSec".
	flagMap[RateSec] = NewFlagStr(RateSec, fmt.Sprint(defValMap[RateSec]))
	// RateMin limits send per minute rate "m", "rateMin".
	flagMap[RateMin] = NewFlagStr(RateMin, fmt.Sprint(defValMap[RateMin]))
	// LogHis enabled log output of histogram, "y", logHis.
	flagMap[LogHis] = NewFlagStr(LogHis, fmt.Sprint(defValMap[LogHis]))
	// ConHis enabled console output of histogram, "z", conHis.
	flagMap[ConHis] = NewFlagStr(ConHis, fmt.Sprint(defValMap[ConHis]))
	// TmoSec request timiout in seconds, "o", tmoSec.
	flagMap[TmoSec] = NewFlagStr(TmoSec, fmt.Sprint(defValMap[TmoSec]))
	// TmoRdS read request timiout in seconds, "i", tmoRdS.
	flagMap[TmoRdS] = NewFlagStr(TmoRdS, fmt.Sprint(defValMap[TmoRdS]))
	// TmoWrS write request timiout in seconds, "j", tmoWrS.
	flagMap[TmoWrS] = NewFlagStr(TmoWrS, fmt.Sprint(defValMap[TmoWrS]))
	// CxnLim limit the number of connecitons, "n", cxnLim.
	flagMap[CxnLim] = NewFlagStr(CxnLim, fmt.Sprint(defValMap[CxnLim]))
	// CltType the type of client, "b", cltType.
	flagMap[CltType] = NewFlagStr(CltType, fmt.Sprint(defValMap[CltType]))
	// SaveReq saves requests, "g", saveReq
	flagMap[SaveReq] = NewFlagStr(SaveReq, fmt.Sprint(defValMap[SaveReq]))
	// SaveRes saves responses, "k", saveRes
	flagMap[SaveRes] = NewFlagStr(SaveRes, fmt.Sprint(defValMap[SaveRes]))
	// TmoCxn connect timeout, "", tmoCxn
	flagMap[TmoCxn] = NewFlagStr(TmoCxn, fmt.Sprint(defValMap[TmoCxn]))

	return flagMap
}

func createExpectedFlags() map[FlagID]*Flag {
	flagMap := createFlags()

	return flagMap
}
