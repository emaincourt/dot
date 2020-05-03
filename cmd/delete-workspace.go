package cmd

import (
	"github.com/spf13/cobra"
)

var deleteWorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Delete a new workspace",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"Deleting workspace %s...",
			args[0],
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.DeleteWorkspace(args[0]); err != nil {
			log.Errorf(
				"An error occured while deleting workspace: %v",
				err,
			)
			return
		}
		log.Info("Successfully deleted workspace")
	},
}

func init() {
	deleteCmd.AddCommand(deleteWorkspaceCmd)
}
