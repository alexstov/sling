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
	"os"
	"path/filepath"
	"strings"

	"github.com/4alexey/sling/conf"
	"github.com/4alexey/sling/cui"
	"github.com/4alexey/sling/slog"
	"github.com/cloudflare/cfssl/log"
	"github.com/kubernetes/pkg/kubectl/util/i18n"
	"github.com/kubernetes/pkg/kubectl/util/templates"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logger slog.Logger
var sconf *conf.SlingConfig

// Con Consoler interface implementation.
var Con cui.Consoler

var (
	rootLong = templates.LongDesc(`
sling is a tool to send files to the endpoint(s),
setting trottling and delays between concurrent requests.`)

	rootExample = templates.Examples(`
		# View current sling config
		sling config view
		
		# Send a single file to specified endpoint
		sling request send -f myfile.dat -a myhost.com -p 9080`)
)

// RootCmd sling root command
var RootCmd = &cobra.Command{
	Use:                   "sling",
	DisableFlagsInUseLine: true,
	Short:                 i18n.T("sling network traffic emulator"),
	Long:                  rootLong,
	Example:               rootExample,
}

// RootFlags command flags
var RootFlags Flagmapper

func init() {
	var err error

	// Create logger with default settings.
	if logger, err = slog.NewLogger(); err != nil {
		fmt.Println("Failed to create logger", err)
		return
	}

	logger.Out(logrus.TraceLevel, nil, "Calling root command initializer")
	if err = initSlingFromConfig(); err != nil {
		fmt.Println("Failed to initialize sling from config.", err)
		return
	}
}

func initSlingFromConfig() (err error) {
	if sconf, err = LoadConfig(os.Getenv("SLINGCONFIG"), logger); err != nil {
		return
	}

	// Set root command flags.
	RootFlags, err = NewCmdFlags(RootCmd, CmdRoot, sconf)

	// Init console
	if Con, err = cui.NewConsole(&sconf.Console, logger); err != nil {
		log.Error("Cannot initialize console.", err)
		return
	}
	
	flagmap := RootFlags.GetFlagmap()

	// The events are used to set flag values passed explicitly from command line.
	flagmap[ConFlat].AddEvent(Events[ConsoleSetFlat], Con.SetFlat)
	flagmap[LogLvl].AddEvent(Events[LogSetLevel], logger.SetLevel)
	flagmap[ConLvl].AddEvent(Events[ConsoleSetLevel], Con.SetLevel)

	if err = logger.Configure(&sconf.Log); err != nil {
		log.Error("Cannot initialize logger.", err)
	}

	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		Con.OutLogAndConsole(logrus.FatalLevel, logrus.Fields{"error": err}, "Root command execution failed.")
		os.Exit(1)
	} else {
		logger.Out(logrus.InfoLevel, nil, "Root command executed.")
	}
}

// LoadConfig loads configuration from the file.
func LoadConfig(path string, logger slog.Logger) (sconf *conf.SlingConfig, err error) {
	logger.Out(logrus.DebugLevel, nil, "Creating new sling config.")
	sconf = conf.NewConfig()

	logger.Out(logrus.DebugLevel, logrus.Fields{"SLINGCONFIG": sconf.Slingconfig}, "Reading sling config.")
	sconf.Slingconfig = path
	viper.SetConfigType("yml")
	basename := filepath.Base(sconf.Slingconfig)
	viper.SetConfigName(strings.TrimSuffix(basename, filepath.Ext(basename)))
	viper.AddConfigPath(filepath.Dir(sconf.Slingconfig))

	if err := viper.ReadInConfig(); err != nil {
		logger.Out(logrus.FatalLevel, logrus.Fields{"SLINGCONFIG": sconf.Slingconfig, "error": err}, "Error reading config file.")
	}

	if err := viper.Unmarshal(sconf); err != nil {
		logger.Out(logrus.FatalLevel, logrus.Fields{"SLINGCONFIG": sconf.Slingconfig, "error": err}, "Error parsing config file.")
	}

	logger.Out(logrus.TraceLevel, logrus.Fields{"SLINGCONFIG": sconf.Slingconfig}, "Successfully created sling config.")

	return sconf, nil
}
