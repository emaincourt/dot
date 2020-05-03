package cmd

import (
	"github.com/spf13/cobra"
)

var createWorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Create a new workspace",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"Creating new workspace %s",
			args[0],
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.CreateWorkspace(args[0]); err != nil {
			log.Errorf(
				"An error occured while creating workspace: %v",
				err,
			)
			return
		}
		log.Info("Successfully created workspace")
	},
}

func init() {
	createCmd.AddCommand(createWorkspaceCmd)
}
