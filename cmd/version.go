package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "Generates and tag git repositories based on semantic versions based and conventional commits",
	Long:  `Generates and tag git repositories based on semantic versions based and conventional commits`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.0") //nolint:forbidigo
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
