package cmd

import (
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:              "use",
	Short:            "Switch use of a resource",
	PersistentPreRun: loadConfig,
}

func init() {
	rootCmd.AddCommand(useCmd)
}
