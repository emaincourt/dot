package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:              "create",
	Short:            "Create a new resource",
	PersistentPreRun: loadConfig,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
