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

	"github.com/alexstov/sling/cmd"
	"github.com/alexstov/sling/unit"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("Root", func() {

	var (
		testCmd *cobra.Command
		flags   cmd.Flagmapper
	)

	BeforeEach(func() {
		var err error

		testCmd = &cobra.Command{
			Use:                   "slingtest",
			DisableFlagsInUseLine: true,
			Short:                 "Root command description",
			Long:                  "Root command log description",
			Example:               "Root command example",
		}
		sconf := unit.NewConfig()

		flags, err = cmd.NewCmdFlags(testCmd, cmd.CmdRoot, &sconf)
		Expect(err).Should(BeNil())
	})

	Describe("Root command", func() {
		Context("created with flags", func() {
			var flagID = cmd.LogLvl
			var flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flags).ShouldNot(BeNil())
				flagmap := flags.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			flagID = cmd.ConLvl
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flags).ShouldNot(BeNil())
				flagmap := flags.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			flagID = cmd.ConFlat
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flags).ShouldNot(BeNil())
				flagmap := flags.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testCmd.Flag(flagName)).ShouldNot(BeNil())
			})
		})
	})
})
