package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:              "sync",
	Short:            "Sync resources with a distant state store",
	PersistentPreRun: loadConfig,
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
