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

package cmd

import (
	"os"
	"os/exec"

	"github.com/kubernetes/pkg/kubectl/util/i18n"
	"github.com/kubernetes/pkg/kubectl/util/templates"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logviewLong = templates.LongDesc(`
		View sling log file`)

	logviewExample = templates.Examples(`
		# View sling log using Linux cat command
		sling log view`)
)

// LogviewCmd command.
var LogviewCmd = &cobra.Command{
	Use:     "view",
	Short:   i18n.T("View sling log file"),
	Long:    logviewLong,
	Example: logviewExample,
	Run:     logviewRun,
}

func logviewRun(cmd *cobra.Command, args []string) {
	viCmd := exec.Command("cat", sconf.Log.LogFile)
	viCmd.Stdout = os.Stdout
	viCmd.Stderr = os.Stderr
	err := viCmd.Run()
	if err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "View command execution failed.")
	} else {
		logger.Out(logrus.InfoLevel, nil, "View command executed.")
	}
}

func init() {
	LogCmd.AddCommand(LogviewCmd)
}
