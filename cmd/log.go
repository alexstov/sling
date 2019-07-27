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
	"github.com/spf13/cobra"
)

// LogCmd command.
var LogCmd = &cobra.Command{
	Use:                   "log SUBCOMMAND",
	DisableFlagsInUseLine: true,
	Short:                 "Manage sling logs",
	Long: `
	Manage sling log file`,
	Example: `
	# View sling log file
	sling log view

	# Clean sling log file
	sling log clean

	# Log and Console levels
	PanicLevel [0]: Highest level of severity
	FatalLevel [1]: Logs and then calls logger.Exit(1).
	ErrorLevel [2]: Used for errors that should definitely be noted.
	WarnLevel  [3]: Non-critical entries that deserve eyes.
	InfoLevel  [4]: General operational entries about what's going on inside the application.
	DebugLevel [5]: Usually only enabled when debugging. Very verbose logging.
	TraceLevel [6]: Designates finer-grained informational events than the Debug.`,
}

func init() {
	RootCmd.AddCommand(LogCmd)
}
