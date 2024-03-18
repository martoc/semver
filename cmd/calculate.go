package cmd

import (
	"fmt"
	"os"

	"github.com/martoc/semver/core"
	"github.com/martoc/semver/logger"
	"github.com/spf13/cobra"
)

func init() {
	calculateCmd.Flags().StringP("path", "p", ".", "Path to a git repository")
}

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculates a new semantic version based on the latest commit message in the repository",
	Long: `Calculates a new semantic version based on the latest commit message in the repository
		using semantic versioning and conventional commits (https://www.conventionalcommits.org/en/v1.0.0-beta.4/)`,
	Run: func(cmd *cobra.Command, _ []string) {
		path, _ := cmd.Flags().GetString("path")
		result, err := core.NewCalculateCommandBuilder().SetPath(path).Build().Execute()
		if err != nil {
			logger.GetInstance().Error(err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, result)
	},
}
