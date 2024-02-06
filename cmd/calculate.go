package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculates a new semantic version based on the latest commit message in the repository",
	Long: `Calculates a new semantic version based on the latest commit message in the repository
		using semantic versioning and conventional commits`,
	Run: func(cmd *cobra.Command, args []string) {
		epoc := time.Now().Unix()
		fmt.Println(epoc) //nolint:forbidigo
	},
}
