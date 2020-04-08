package cmd

import (
	"foundry/cli/logger"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "init",
		Short:   "Create the initial foundry.yaml config file",
		Example: "foundry init",
		Run:     runInit,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	if _, err := os.Stat(confFile); !os.IsNotExist(err) {
		logger.FdebuglnError("Foundry config file 'foundry.yaml' already exists")
		logger.FatalLogln("Foundry config file 'foundry.yaml' already exists")
	}

	dest := filepath.Join(foundryConf.CurrentDir, "foundry.yaml")
	err := ioutil.WriteFile(dest, []byte(getInitYaml()), 0644)
	if err != nil {
		logger.FdebuglnError("Error writing foundry.yaml:", err)
		logger.FatalLogln("Error creating Foundry config file 'foundry.yaml':", err)
	}

	logger.SuccessLogln("Config file 'foundry.yaml' created")
}

func getInitYaml() string {
	// TODO: Update to a final version of the init config yaml
	return `
# An array of glob patterns for files that should be ignored. The path is relative to the root dir.
# If the array is changed, the CLI must be restarted for it to take the effect
ignore:
  - node_modules # Skip the whole node_modules directory
  - .git # Skip the whole .git directory
  - "**/.*" # Skip all hidden files
  - "**/*~" # Skip vim's temp files
# An array of Firebase functions that should be evaluated by Foundry. All these functions must be exported in your root index.js
functions:
#   - name: hello
#     type: https
#     payload: '{"key":"world"}'
#   - name: hello
#     type: https
#     payload: '{"key":"value2"}'
`
}
