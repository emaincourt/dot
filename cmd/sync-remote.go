package cmd

import (
	"github.com/spf13/cobra"
)

var syncRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Add/Delete remotes",
}

func init() {
	syncCmd.AddCommand(syncRemoteCmd)
}
