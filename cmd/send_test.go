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

	"github.com/4alexey/sling/cmd"
	"github.com/4alexey/sling/emul"
	"github.com/4alexey/sling/unit"
	"github.com/kubernetes/pkg/kubectl/util/i18n"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

var _ = Describe("Send", func() {

	var (
		testSendCmd *cobra.Command
		flagmapper  cmd.Flagmapper
	)

	BeforeEach(func() {
		var err error
		testSendCmd = &cobra.Command{
			Use:   "send [flags]",
			Short: i18n.T("Send requests to the endpoint"),
			Long: `Send requests from a file or directory to the endpoint.

			Examples:
			# Send myfile.dat from /tmp/data directory. Use default send parameters from SLINGCONFIG.
			sling request send -f myfile.dat -d /tmp/data
			`,
		}

		sconf := unit.NewConfig()

		flagmapper, err = cmd.NewCmdFlags(testSendCmd, cmd.CmdSend, &sconf)
		Expect(err).Should(BeNil())
	})

	Describe("Send command", func() {
		Context("created with flags", func() {
			// file name matching wildcard, -w, --wildcard
			var flagID = cmd.Wildcard
			var flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// send repeat count, -r, -repeat
			flagID = cmd.Repeat
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// directory to send files from, -d, --dir
			flagID = cmd.Dir
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// number of concurrent connections, -n, --cxnNum
			flagID = cmd.CxnNum
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// delay after each repeated request, -e, --sleepMs
			flagID = cmd.SleepMs
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// directory to save requests, -k. --saveReqDir
			flagID = cmd.SaveReqDir
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// directory to save response, -j, --saveReqDir
			flagID = cmd.SaveResDir
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// endpoint IP, DNS name, or HTTP address, -a, --address
			flagID = cmd.Address
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// endpoint port number, -p, --port
			flagID = cmd.Port
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// send rate per second, -s, --reteSec
			flagID = cmd.RateSec
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// send rate per minute, -m, --rateMin
			flagID = cmd.RateMin
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// write histogram to log file, -g, --logHis
			flagID = cmd.LogHis
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// write histogram to console, -y, --conHis
			flagID = cmd.ConHis
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// network client type, TCP or HttpPost, -c, --cltType
			flagID = cmd.CltType
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// limit the number of concurrent connections, -l, --cxnLim
			flagID = cmd.CxnLim
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// active endpoint index in SLINGCONFIG, zero-based, -i, --endpoint
			flagID = cmd.Endpoint
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// file path or file name to send, -d, --file
			flagID = cmd.File
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// save requests, -q, --saveReq
			flagID = cmd.SaveReq
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// save responses, -o, --saveRes
			flagID = cmd.SaveRes
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// network client dial timeout, -u, --tmoCxn
			flagID = cmd.TmoCxn
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// read network client timeout, v, --tmoRdS
			flagID = cmd.TmoRdS
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// network client timeout, -t, --tmoSec
			flagID = cmd.TmoSec
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
			// write network client timeout, -x, tmoWrS
			flagID = cmd.TmoWrS
			flagName = fmt.Sprintf("%s", flagID)
			It(flagName, func() {
				defer GinkgoRecover()
				Expect(flagmapper).ShouldNot(BeNil())
				flagmap := flagmapper.GetFlagmap()
				Expect(flagmap[flagID]).ShouldNot(BeNil())
				Expect(testSendCmd.Flag(flagName)).ShouldNot(BeNil())
			})
		})
	})

	Describe("Send command", func() {
		Context("NewSendEmul", func() {
			It("created with valid SendArgs", func() {
				defer GinkgoRecover()
				var err error
				var em *emul.Emul
				sendArgs := unit.NewSendArgs()
				em, err = cmd.NewSendEmul(&sendArgs)
				Expect(em).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})
		})
	})
})
