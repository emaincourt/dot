package cmd

import (
	"github.com/spf13/cobra"
)

var useWorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Change the current workspace",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof("Switching active workspace to %s...", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.SetActiveWorkspace(args[0]); err != nil {
			log.Errorf(
				"An error occured while setting new active workspace: %v",
				err,
			)
			return
		}

		log.Info("Successfully modified active workspace")
	},
}

func init() {
	useCmd.AddCommand(useWorkspaceCmd)
}
