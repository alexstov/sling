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
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// ViewCmd command.
var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View slingconfig settings",
	Long: `
	View slingconfig settings`,
	Example: `
	# View current slingconfig settings
	sling config view`,
	Run: viewRun,
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]

	return ok
}

func viewRun(cmd *cobra.Command, args []string) {
	err := PrintObj(sconf, os.Stdout)

	if err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "View command execution failed.")
	} else {
		logger.Out(logrus.InfoLevel, nil, "View command executed.")
	}
}

func init() {
	ConfigCmd.AddCommand(ViewCmd)
}

// PrintObj impl
func PrintObj(obj interface{}, w io.Writer) error {
	output, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, string(output))

	return err
}
