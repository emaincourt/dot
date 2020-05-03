package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var getEnv = &cobra.Command{
	Use:     "envs",
	Aliases: []string{"env"},
	Short:   "List the environment variables for the current workspace",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Listing Env...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		workspace, err := cfg.GetActiveWorkspace()
		if err != nil {
			log.Errorf(
				"An error occured while listing Env: %v",
				err,
			)
			return
		}

		data := make([][]string, len(workspace.Env))
		for _, env := range workspace.Env {
			data = append(data, []string{
				env.Name,
				env.Value,
				strconv.FormatBool(env.Private),
			})
		}
		log.Table(
			[]string{"Name", "Value", "Private"},
			data,
		)
	},
}

func init() {
	getCmd.AddCommand(getEnv)
}
