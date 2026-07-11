package cmd

import "github.com/spf13/cobra"

var discoverCmd = &cobra.Command{
    Use:   "discover",
    Short: "Discover projects in one or more workspaces",
    RunE: func(cmd *cobra.Command, args []string) error {

        return nil
    },
}

func init() {
    workspaceCmd.AddCommand(discoverCmd)
}