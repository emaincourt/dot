package viper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/emaincourt/dot/pkg/config"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func (c *ViperConfig) CreateWorkspace(name string) error {
	workspaces := c.root.GetStringMapString(config.RootConfigWorkspacesKey)
	workspaces[name] = fmt.Sprintf(
		"%s/%s.%s",
		config.DefaultWorkspacesDir,
		name,
		config.DefaultConfigFileType,
	)

	c.root.Set(
		config.RootConfigWorkspacesKey,
		workspaces,
	)

	path := path.Join(
		path.Dir(c.root.ConfigFileUsed()),
		workspaces[name],
	)

	if _, err := os.Create(path); err != nil {
		return c.Error(err)
	}

	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return c.Error(errors.New(ErrNoRead))
	}

	v.Set(config.ActiveWorkspaceConfigSourcesKey, []config.WorkspaceSource{})
	v.Set(config.ActiveWorkspaceEnvKey, []config.WorkspaceEnv{})

	errGroup, _ := errgroup.WithContext(context.Background())
	errGroup.Go(func() error {
		return c.Write(
			WithViper(c.root),
		)
	})
	errGroup.Go(func() error {
		return c.Write(
			WithViper(v),
		)
	})

	return errGroup.Wait()
}

func (c *ViperConfig) DeleteWorkspace(name string) error {
	workspaces, err := c.GetWorkspaces()
	if err != nil {
		return c.Error(err)
	}

	workspace, ok := workspaces[name]
	if ok {
		if c.root.GetString(config.RootConfigActiveKey) == name {
			return c.Error(errors.New(ErrWorkspaceInUse))
		}

		delete(workspaces, name)
		c.root.Set(config.RootConfigWorkspacesKey, workspaces)
	} else {
		workspace = fmt.Sprintf(
			"%s/%s%s",
			config.DefaultWorkspacesDir,
			name,
			config.DefaultConfigFileExtension,
		)
	}

	path := path.Join(
		path.Dir(c.root.ConfigFileUsed()),
		workspace,
	)

	_ = os.Remove(path)

	return c.Write(WithViper(c.root))
}

func (c *ViperConfig) GetActiveWorkspace() (*config.WorkspaceConfig, error) {
	var cfg config.WorkspaceConfig
	if err := c.active.Unmarshal(&cfg); err != nil {
		return nil, c.Error(
			err,
			WithViper(c.active),
		)
	}

	return &cfg, nil
}

func (c *ViperConfig) SetActiveWorkspace(name string) error {
	workspaces := c.root.GetStringMapString(config.RootConfigWorkspacesKey)

	var exists bool
	for workspace := range workspaces {
		if workspace == name {
			exists = true
		}
	}

	if !exists {
		return c.Error(errors.New(ErrUnknownWorkspace))
	}

	c.root.Set(
		config.RootConfigActiveKey,
		name,
	)

	return c.Write(WithViper(c.root))
}

func (c *ViperConfig) GetWorkspaces() (map[string]string, error) {
	if !c.root.IsSet(config.RootConfigWorkspacesKey) {
		return nil, c.Error(errors.New(ErrNotSet))
	}

	return c.root.GetStringMapString(config.RootConfigWorkspacesKey), nil
}
