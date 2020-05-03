package cmd

import (
	"github.com/spf13/cobra"
)

var UnsetCmd = &cobra.Command{
	Use:              "unset",
	Short:            "Unset an environment variable for current workspace",
	Args:             cobra.MaximumNArgs(1),
	PersistentPreRun: loadConfig,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof("Removing env var %s from active workspace...", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.UnSetEnvironmentVariableForActiveWorkspace(args[0]); err != nil {
			log.Errorf(
				"An error occured while unsetting env var: %v",
				err,
			)
			return
		}

		log.Info("Successfully removed env var")
	},
}

func init() {
	rootCmd.AddCommand(UnsetCmd)
}
