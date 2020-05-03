package viper

import (
	"errors"

	"github.com/emaincourt/dot/pkg/config"
)

func (c *ViperConfig) GetSourcesForActiveWorkspace() ([]config.WorkspaceSource, error) {
	if !c.active.IsSet(config.ActiveWorkspaceConfigSourcesKey) {
		return nil, c.Error(errors.New(ErrNotSet))
	}

	return c.active.GetStringSlice(config.ActiveWorkspaceConfigSourcesKey), nil
}

func (c *ViperConfig) AddSourceForActiveWorkspace(source string) error {
	if !c.active.IsSet(config.ActiveWorkspaceConfigSourcesKey) {
		return c.Error(errors.New(ErrNotSet))
	}

	c.active.Set(
		config.ActiveWorkspaceConfigSourcesKey,
		append(
			c.active.GetStringSlice(config.ActiveWorkspaceConfigSourcesKey),
			source,
		),
	)

	return c.Write(WithViper(c.active))
}
