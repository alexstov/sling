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
	// "github.com/alexstov/sling/sio"
	"os"
	"os/exec"

	"github.com/alexstov/sling/sio"
	"github.com/kubernetes/pkg/kubectl/util/i18n"
	"github.com/kubernetes/pkg/kubectl/util/templates"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logcleanLong = templates.LongDesc(`
		Clean sling log file`)

	logcleanExample = templates.Examples(`
		# Clean sling log file using Linux truncate command
		sling log clean`)
)

// LogcleanCmd command.
var LogcleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   i18n.T("Clean sling log file"),
	Long:    logcleanLong,
	Example: logcleanExample,
	Run:     logcleanRun,
}

func logcleanRun(cmd *cobra.Command, args []string) {
	var err error
	var filer sio.Filer

	/*
		Clean log file.
	*/
	viCmd := exec.Command("truncate", "-s", "0", sconf.Log.LogFile)

	viCmd.Stdout = os.Stdout
	viCmd.Stderr = os.Stderr
	err = viCmd.Run()

	if err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "Clean command execution failed.")
		return
	} /**/

	// Create command filer.
	if filer, err = sio.NewFiler(logger); err != nil {
		logger.Out(logrus.ErrorLevel, logrus.Fields{"err": err}, "Cannot create filer.")
	}

	/*
		Clean log directories.
	*/
	if err = filer.CleanDir(sconf.SaveReqDir); err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "Clean command execution failed.")
		return
	}
	if err = filer.CleanDir(sconf.SaveResDir); err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "Clean command execution failed.")
		return
	} /**/

	Con.OutLogAndConsole(logrus.InfoLevel, nil, "Logs cleaned.")
}

func init() {
	LogCmd.AddCommand(LogcleanCmd)
}
