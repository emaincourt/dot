package cmd

import (
	"github.com/spf13/cobra"
)

var syncRemoteAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new remote",
	Args:  cobra.MinimumNArgs(2),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Adding new remote...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := state.AddRemote(args[0], args[1]); err != nil {
			log.Errorf(
				"An error occured while adding remote: %v",
				err,
			)
			return
		}

		log.Info("Successfully added remote")
	},
}

func init() {
	syncRemoteCmd.AddCommand(syncRemoteAddCmd)
}
