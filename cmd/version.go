package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/martoc/semver/logger"
	"github.com/spf13/cobra"
)

var (
	CLIVersion string
	Os         string //nolint:varnamelen
	Arch       string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions`,
	Run: func(_ *cobra.Command, _ []string) {
		jsonData := struct {
			Version string `json:"version"`
			Os      string `json:"os"`
			Arch    string `json:"arch"`
		}{
			Version: CLIVersion,
			Os:      Os,
			Arch:    Arch,
		}
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			logger.GetInstance().Println("Error marshaling version data:", err)
		}
		fmt.Fprintln(os.Stdout, string(jsonBytes))
	},
}
