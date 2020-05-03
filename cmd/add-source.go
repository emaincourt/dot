package cmd

import (
	"github.com/spf13/cobra"
)

var addSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Appends a new source to the list of existing ones",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof(
			"Adding new source %s...",
			args[0],
		)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.AddSourceForActiveWorkspace(
			args[0],
		); err != nil {
			log.Errorf(
				"An error occured while adding source: %v",
				err,
			)
			return
		}
		log.Info("Successfully added source")
	},
}

func init() {
	addCmd.AddCommand(addSourceCmd)
}
