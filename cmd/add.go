package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:              "add",
	Short:            "Appends a new value to a list of existing values",
	PersistentPreRun: loadConfig,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
