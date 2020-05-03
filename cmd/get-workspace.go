package cmd

import (
	"github.com/spf13/cobra"
)

var getWorkspaces = &cobra.Command{
	Use:     "workspaces",
	Aliases: []string{"workspace"},
	Short:   "List the available workspaces",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Listing workspaces...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		workspaces, err := cfg.GetWorkspaces()
		if err != nil {
			log.Errorf(
				"An error occured while listing workspaces: %v",
				err,
			)
			return
		}

		data := make([][]string, len(workspaces))
		for _, workspace := range workspaces {
			data = append(data, []string{
				workspace,
			})
		}
		log.Table(
			[]string{"Workspaces"},
			data,
		)
	},
}

func init() {
	getCmd.AddCommand(getWorkspaces)
}
