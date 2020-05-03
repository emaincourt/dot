package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:              "get",
	Short:            "List a type of resources",
	PersistentPreRun: loadConfig,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
