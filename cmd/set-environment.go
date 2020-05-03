package cmd

import (
	"github.com/spf13/cobra"
)

const (
	FlagPrivate      = "private"
	FlagPrivateShort = "p"
)

var (
	private bool
)

var setCmd = &cobra.Command{
	Use:              "set",
	Short:            "Set a new environment variable for current workspace",
	Args:             cobra.MinimumNArgs(2),
	PersistentPreRun: loadConfig,
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Infof("Setting new env var %s...", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.SetEnvironmentVariableForActiveWorkspace(
			args[0],
			args[1],
			private,
		); err != nil {
			log.Errorf(
				"An error occured while setting new env var: %v",
				err,
			)
			return
		}

		log.Info("Successfully set env var")
	},
}

func init() {
	setCmd.Flags().BoolVarP(&private, FlagPrivate, FlagPrivateShort, false, "This flag should be set to true is the environment variable should not be committed on sync")
	rootCmd.AddCommand(setCmd)
}
