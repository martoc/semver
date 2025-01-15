package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/martoc/semver/core"
	"github.com/martoc/semver/logger"
	"github.com/spf13/cobra"
)

func init() {
	calculateCmd.Flags().StringP("path", "p", ".", "Path to a git repository")
	calculateCmd.Flags().BoolP("push", "u", false, "Push the new tag to the remote repository")
	calculateCmd.Flags().BoolP("add-floating-tags", "f", false,
		"Add the floating tags to the new tag for example v1.2.3 will also add v1 and v1.2")
	calculateCmd.Flags().BoolP("disable-tagging", "d", false, "Disable tagging")
}

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculates a new semantic version based on the latest commit message in the repository",
	Long: `Calculates a new semantic version based on the latest commit message in the repository
		using semantic versioning and conventional commits (https://www.conventionalcommits.org/en/v1.0.0-beta.4/)`,
	Run: func(cmd *cobra.Command, _ []string) {
		path, _ := cmd.Flags().GetString("path")
		push, _ := cmd.Flags().GetBool("push")
		addFloatingTags, _ := cmd.Flags().GetBool("add-floating-tags")
		disableTagging, _ := cmd.Flags().GetBool("disable-tagging")
		result, err := core.NewCalculateCommandBuilder().
			SetPath(path).
			SetAddFloatingTags(addFloatingTags).
			SetPush(push).
			SetDisableTagging(disableTagging).
			Build().
			Execute()
		if err != nil {
			logger.GetInstance().Error(err)
			os.Exit(1)
		}
		jsonResult, err := json.Marshal(result) // Convert result to JSON
		if err != nil {
			logger.GetInstance().Error(err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, string(jsonResult)) // Print JSON result
	},
}
