package cmd

import (
	"github.com/spf13/cobra"
)

var tidyCmd = &cobra.Command{
	Use:              "tidy",
	Short:            "Clean up workspaces",
	Args:             cobra.MaximumNArgs(1),
	PersistentPreRun: loadConfig,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Cleaning up configuration...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.Tidy(); err != nil {
			log.Errorf(
				"An error occured while cleaning up configuration: %v",
				err,
			)
			return
		}

		log.Info("Successfully cleaned up configuration")
	},
}

func init() {
	rootCmd.AddCommand(tidyCmd)
}
