package cmd

import (
	"os"

	"github.com/martoc/semver/core"
	"github.com/martoc/semver/logger"
	"github.com/spf13/cobra"
)

func init() {
	calculateCmd.Flags().StringP("path", "p", ".", "Path to the git repository")
}

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculates a new semantic version based on the latest commit message in the repository",
	Long: `Calculates a new semantic version based on the latest commit message in the repository
		using semantic versioning and conventional commits`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		result, err := core.NewCalculateCommandBuilder().SetPath(path).Build().Execute()
		if err != nil {
			logger.Instance.Println(err)
			os.Exit(1)
		}
		cmd.Println(result)
	},
}
