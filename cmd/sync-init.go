package cmd

import (
	"github.com/spf13/cobra"
)

var syncInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new remote state",
	Run: func(cmd *cobra.Command, args []string) {
		if err := state.Init(); err != nil {
			panic(err)
		}
	},
}

func init() {
	syncCmd.AddCommand(syncInitCmd)
}
