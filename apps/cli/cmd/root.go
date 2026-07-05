package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ruach",
	Short: "An AI-native personal engineering platform",
	Long: `Ruach is an AI-native personal engineering platform
designed to help developers build, learn, automate,
and remember.`,
}

func Execute() error {
	return rootCmd.Execute()
}
