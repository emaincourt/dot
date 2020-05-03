package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/emaincourt/dot/pkg/config"
	"github.com/emaincourt/dot/pkg/config/viper"
	"github.com/emaincourt/dot/pkg/crypto"
	aescrypto "github.com/emaincourt/dot/pkg/crypto/aes"
	"github.com/emaincourt/dot/pkg/logger"
	fmtlogger "github.com/emaincourt/dot/pkg/logger/fmt"
	"github.com/emaincourt/dot/pkg/rc"
	shrc "github.com/emaincourt/dot/pkg/rc/sh"
	"github.com/emaincourt/dot/pkg/sync"
	"github.com/emaincourt/dot/pkg/sync/git"
	"github.com/emaincourt/dot/pkg/zip"
	"github.com/emaincourt/dot/pkg/zip/archiver"

	"github.com/spf13/cobra"
)

var (
	cfgPath          string
	encryptionSecret string
	logLevel         string
	noRegenerate     bool
	rcFilePath       string

	cfg         config.Config
	state       sync.Sync
	encryptor   crypto.Crypto
	zipper      zip.Zipper
	log         logger.Logger
	regenerator rc.RCGenerator
)

var rootCmd = &cobra.Command{
	Use: "dot",
	Run: func(c *cobra.Command, args []string) {},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if err := regenerator.Regenerate(rcFilePath); err != nil {
				log.Errorf(
					"An error occured while regenerating rc file: %v",
					err,
				)
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf(
			"An error occured while executing command: %v",
			err,
		)
		os.Exit(1)
	}
}

const (
	FlagConfig             = "config"
	FlagConfigShort        = "c"
	FlagEncryptionKey      = "encryption-key"
	FlagEncryptionKeyShort = "e"
	FlagLogLevel           = "log-level"
	FlagRCFilePath         = "rc-file-path"
	FlagRCFilePathShort    = "f"
	FlagNoRegenerate       = "no-regenerate"
	FlagNoRegenerateShort  = "n"

	DefaultConfigPath    = "$HOME/.dot/dot.yaml"
	DefaultEncryptionKey = ""
	DefaultNoRegenerate  = false
	DefaultRCFilePath    = "$HOME/.dot/.dotrc"

	ConfigNameDot       = "dot"
	ConfigNameWorkspace = "workspace"

	WorkspacesKey      = "workspaces"
	ActiveWorkspaceKey = "active"

	WorkspaceSourcesKey = "sources"
	WorkspaceEnvKey     = "env"
)

func init() {
	cobra.OnInitialize(onInitialize)
	rootCmd.PersistentFlags().StringVarP(
		&cfgPath,
		FlagConfig,
		FlagConfigShort,
		strings.Replace(DefaultConfigPath, "$HOME", os.Getenv("HOME"), 1),
		"Path to the dot configuration file",
	)
	rootCmd.PersistentFlags().StringVarP(
		&encryptionSecret,
		FlagEncryptionKey,
		FlagEncryptionKeyShort,
		DefaultEncryptionKey,
		"The secret to use to encrypt/decrypt data on sync",
	)
	rootCmd.PersistentFlags().StringVar(
		&logLevel,
		FlagLogLevel,
		logger.DefaultLogLevel,
		"The verbosity of output: debug, info, warn, error",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&noRegenerate,
		FlagNoRegenerate,
		FlagNoRegenerateShort,
		DefaultNoRegenerate,
		"Whether or not the rc file should be regenerated after any operation",
	)
	rootCmd.PersistentFlags().StringVarP(
		&rcFilePath,
		FlagRCFilePath,
		FlagRCFilePathShort,
		strings.Replace(DefaultRCFilePath, "$HOME", os.Getenv("HOME"), 1),
		"The path to the rc file to regenerate",
	)
}

func onInitialize() {
	log = fmtlogger.NewFmtLogger(
		logLevel,
	)

	cfgPath, _ = rootCmd.Flags().GetString(FlagConfig)

	cfg = viper.NewViperConfig()
	state = git.NewGitSync(path.Dir(cfgPath))
	encryptor = aescrypto.NewAESCrypto(encryptionSecret)
	zipper = archiver.NewArchiver()
	regenerator = shrc.NewShRCGenerator(cfg)
}

func loadConfig(cmd *cobra.Command, args []string) {
	cfgPath, _ = rootCmd.Flags().GetString(FlagConfig)

	if err := cfg.Load(cfgPath); err != nil {
		log.Error("Could not load config file, did you run `dot init` ?")
		os.Exit(0)
	}
}
