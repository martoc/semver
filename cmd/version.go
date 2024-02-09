package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CLIVersion string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stdout, CLIVersion)
	},
}
