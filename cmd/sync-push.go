package cmd

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	FlagCompressTarget      = "compress-target"
	FlagCompressTargetShort = "t"
	FlagCommitMessage       = "commit-message"
	FlagCommitMessageShort  = "m"

	DefaultCompressTarget = "dot.zip"
)

var (
	compressTarget string
	commitMessage  string
)

var syncPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Updates the remote state with local changes",
	PreRun: func(cmd *cobra.Command, args []string) {
		log.Info("Updating the remote state with local changes...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		files := []string{}
		if err := filepath.Walk(path.Dir(cfgPath),
			func(p string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && !strings.Contains(p, ".git") {
					files = append(files, p)
				}
				return nil
			}); err != nil {
			log.Errorf(
				"An error occured while listing files: %v",
				err,
			)
			return
		}

		for i, file := range files {
			fileName, err := encryptor.Encrypt(
				file,
			)
			if err != nil {
				log.Errorf(
					"An error occured while encrypting file %s: %v",
					file,
					err,
				)
				return
			}
			files[i] = fileName
			defer os.Remove(fileName)
		}

		if err := state.Push(
			commitMessage,
			files,
		); err != nil {
			log.Errorf(
				"An error occured while pushing state: %v",
				err,
			)
			return
		}

		log.Info("Successfully pushed local changes")
	},
}

func init() {
	syncPushCmd.Flags().StringVarP(&compressTarget, FlagCompressTarget, FlagCompressTargetShort, DefaultCompressTarget, "The name of the encrypted bundle")
	syncPushCmd.Flags().StringVarP(&commitMessage, FlagCommitMessage, FlagCommitMessageShort, "", "The message to attach to the push (e.g. git commit message if usig git as state store)")
	syncCmd.AddCommand(syncPushCmd)
}
