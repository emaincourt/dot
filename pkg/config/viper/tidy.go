package viper

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/emaincourt/dot/pkg/config"
)

func (c *ViperConfig) Tidy() error {
	var workspaces map[config.WorkspaceName]string
	if err := c.root.UnmarshalKey(config.RootConfigWorkspacesKey, &workspaces); err != nil {
		return c.Error(err)
	}

	workspacesCfgs, err := ioutil.ReadDir(
		path.Join(
			path.Dir(c.root.ConfigFileUsed()),
			config.DefaultWorkspacesDir,
		),
	)
	if err != nil {
		return c.Error(err)
	}

	for workspace := range workspaces {
		var hasConfigFile bool
		for _, cfg := range workspacesCfgs {
			fileName := strings.TrimSuffix(
				cfg.Name(),
				config.DefaultConfigFileExtension,
			)
			if fileName == workspace {
				hasConfigFile = true
			}
		}

		if !hasConfigFile {
			if err := c.CreateWorkspace(workspace); err != nil {
				return c.Error(err)
			}
		}
	}

	for _, cfg := range workspacesCfgs {
		fileName := strings.TrimSuffix(
			cfg.Name(),
			config.DefaultConfigFileExtension,
		)

		var hasWorkspace bool
		for workspace := range workspaces {
			if fileName == workspace {
				hasWorkspace = true
			}
		}

		if !hasWorkspace {
			if err := c.DeleteWorkspace(fileName); err != nil {
				return c.Error(err)
			}
		}
	}

	return nil
}
