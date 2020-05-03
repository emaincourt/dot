package cmd

import (
	"github.com/spf13/cobra"
)

const (
	DefaultDotDirPerm = 0644
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialises dot by creating the configuration folder and associated files",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Initializing dot...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.Init(cfgPath); err != nil {
			log.Errorf(
				"An error occured while initializing dot: %v",
				err,
			)
			return
		}

		log.Info("Successfully initialized dot")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
