package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(calculateCmd)
}

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "Generates and tag git repositories based on semantic versions based and conventional commits",
	Long:  `Generates and tag git repositories based on semantic versions based and conventional commits`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
