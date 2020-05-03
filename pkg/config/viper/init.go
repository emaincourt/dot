package viper

import (
	"os"
	"path"

	"github.com/emaincourt/dot/pkg/config"
	"github.com/spf13/viper"
)

const (
	DefaultDotDirPerm = 0744
)

func (c *ViperConfig) Init(cfgPath string) error {
	if _, err := os.Stat(path.Dir(cfgPath)); err != nil {
		if err := os.MkdirAll(
			path.Dir(cfgPath),
			DefaultDotDirPerm,
		); err != nil {
			return c.Error(err)
		}
	}

	workspacesDir := path.Join(
		path.Dir(cfgPath),
		config.DefaultWorkspacesDir,
	)
	if _, err := os.Stat(workspacesDir); err != nil {
		if err := os.MkdirAll(
			workspacesDir,
			DefaultDotDirPerm,
		); err != nil {
			return c.Error(err)
		}
	}

	if _, err := os.Stat(cfgPath); err != nil {
		if _, err := os.Create(
			cfgPath,
		); err != nil {
			return c.Error(err)
		}
	}

	viper.SetConfigFile(cfgPath)
	viper.Set(config.RootConfigActiveKey, "")
	viper.Set(config.RootConfigWorkspacesKey, make(map[config.WorkspaceName]string))

	if err := viper.WriteConfig(); err != nil {
		return c.Error(err)
	}
	return nil
}
