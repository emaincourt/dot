package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:              "delete",
	Short:            "Deletes a resources which exists or not",
	PersistentPreRun: loadConfig,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
