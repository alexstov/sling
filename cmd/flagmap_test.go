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
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"

	. "github.com/alexstov/sling/cmd"
	"github.com/alexstov/sling/unit"
)

var _ = Describe("Flagmap", func() {

	var (
		err      error
		t        testing.T
		mockCtrl *gomock.Controller
		// ctx         context.Context
		testCmd *cobra.Command
		flags   Flagmapper
		flag    *Flag
		// pflag   *pflag.Flag
	)

	BeforeEach(func() {
		// ctx = context.TODO()
		mockCtrl = gomock.NewController(&t)

		testCmd = &cobra.Command{
			Use:                   "slingtest",
			DisableFlagsInUseLine: true,
			Short:                 "root cmd long descr",
			Long:                  "root cmd long descr",
			Example:               "root cmd example",
		}
	})

	Describe("NewCmdFlags", func() {
		Context("root command", func() {
			It("flags created", func() {
				defer GinkgoRecover()
				defer mockCtrl.Finish()
				sconf := unit.NewConfig()

				flags, err = NewCmdFlags(testCmd, CmdRoot, &sconf)

				Expect(err).Should(BeNil())
				Expect(flags).ShouldNot(BeNil())
				pflags := testCmd.Flags()
				Expect(pflags).ShouldNot(BeNil())
				flagMap := flags.GetFlagmap()
				Expect(len(flagMap)).To(Equal(3))
				flag = flagMap[LogLvl]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[ConLvl]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[ConFlat]
				Expect(flag).ShouldNot(BeNil())

				// pflag = pflags.Lookup(fmt.Sprintf("%s", LogLvl))
				// Expect(pflag).ShouldNot(BeNil())
				// pflag = pflags.Lookup(fmt.Sprintf("%s", ConLvl))
				// Expect(pflag).ShouldNot(BeNil())
				// pflag = pflags.Lookup(fmt.Sprintf("%s", ConFlat))
				// Expect(pflag).ShouldNot(BeNil())
			})
		})

		Context("send command", func() {
			It("flags created", func() {
				defer GinkgoRecover()
				defer mockCtrl.Finish()
				sconf := unit.NewConfig()

				flags, err = NewCmdFlags(testCmd, CmdSend, &sconf)

				Expect(err).Should(BeNil())
				Expect(flags).ShouldNot(BeNil())
				pflags := testCmd.Flags()
				Expect(pflags).ShouldNot(BeNil())
				flagMap := flags.GetFlagmap()
				Expect(23).To(Equal(len(flagMap)))
				flag = flagMap[File]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[Repeat]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[Dir]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[Wildcard]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[CxnNum]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[SleepMs]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[SaveReqDir]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[SaveResDir]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[Address]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[Port]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[RateSec]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[RateMin]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[LogHis]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[ConHis]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[TmoSec]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[TmoRdS]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[TmoWrS]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[CxnLim]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[CltType]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[SaveReq]
				Expect(flag).ShouldNot(BeNil())
				flag = flagMap[SaveRes]
				Expect(flag).ShouldNot(BeNil())
			})
		})
	})
})
