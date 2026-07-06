package cmd

import (
	"fmt"

	"github.com/Jedidiah-Show/ruach/apps/cli/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ruach",
	Short: "An AI-native personal engineering platform",
	Long:  `Ruach is an AI-native personal engineering platform`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no flags or subcommands are passed, show the help menu
		cmd.Help()
	},
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate the autocompletion script for the specified shell",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`Ruach

`)

		fmt.Printf("Version	: %s\n", version.Version)
		fmt.Printf("Commit	: %s\n", version.Commit)
		fmt.Printf("Built	: %s\n", version.Date)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
