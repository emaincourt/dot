package cmd

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var syncPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls changes from the remote state",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Pulling changes from the remote state...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := state.Pull(); err != nil {
			log.Errorf(
				"An error occured while pulling remote state: %v",
				err,
			)
			return
		}

		encryptedFiles := []string{}
		if err := filepath.Walk(path.Dir(cfgPath),
			func(p string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.Contains(p, ".encrypted") {
					encryptedFiles = append(encryptedFiles, p)
				}
				return nil
			}); err != nil {
			log.Errorf(
				"An error occured while listing files: %v",
				err,
			)
			return
		}

		for _, file := range encryptedFiles {
			err := encryptor.Decrypt(
				file,
			)
			if err != nil {
				log.Errorf(
					"An error occured while decrypting file %s: %v",
					file,
					err,
				)
				return
			}
			defer os.Remove(file)
		}

		log.Info("Successfully pulled local changes")
	},
}

func init() {
	syncCmd.AddCommand(syncPullCmd)
}
