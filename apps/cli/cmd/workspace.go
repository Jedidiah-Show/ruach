package cmd

import "github.com/spf13/cobra"

var workspaceCmd = &cobra.Command{
    Use:   "workspace",
    Short: "Discover projects in one or more workspaces",
    RunE: func(cmd *cobra.Command, args []string) error {

        return nil
    },
}

func init() {
    rootCmd.AddCommand(workspaceCmd)
}