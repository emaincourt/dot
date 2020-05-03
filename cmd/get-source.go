package cmd

import (
	"github.com/spf13/cobra"
)

var getSources = &cobra.Command{
	Use:     "sources",
	Aliases: []string{"source"},
	Short:   "List the sources for the current workspace",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Listing sources...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		sources, err := cfg.GetSourcesForActiveWorkspace()
		if err != nil {
			log.Errorf(
				"An error occured while listing sources: %v",
				err,
			)
			return
		}

		data := make([][]string, len(sources))
		for _, source := range sources {
			data = append(data, []string{source})
		}
		log.Table(
			[]string{"Sources"},
			data,
		)
	},
}

func init() {
	getCmd.AddCommand(getSources)
}
