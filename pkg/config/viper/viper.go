package viper

import (
	"fmt"
	"path"

	"github.com/emaincourt/dot/pkg/config"
	"github.com/spf13/viper"
)

const (
	ErrNotSet           = "missing value in config"
	ErrUnknownWorkspace = "unknown workspace"
	ErrNoRead           = "cannot read from config file"
	ErrWorkspaceInUse   = "cannot delete a context which is currently active"
)

type ViperConfig struct {
	root   *viper.Viper
	active *viper.Viper
}

func (c *ViperConfig) Error(
	err error,
	o ...Option,
) error {
	opts := &Options{}
	for _, opt := range o {
		opt(opts)
	}

	v := viper.GetViper()
	if opts.viper != nil {
		v = opts.viper
	}

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		return fmt.Errorf("Could not find configuration file for path %s", v.ConfigFileUsed())
	} else {
		return err
	}
}

func NewViperConfig() *ViperConfig {
	return &ViperConfig{
		root:   viper.New(),
		active: viper.New(),
	}
}

func (c *ViperConfig) Load(rootCfgPath string) error {
	c.root.SetConfigFile(rootCfgPath)
	if err := c.root.ReadInConfig(); err != nil {
		return c.Error(err)
	}

	activeWorkspace := c.root.GetString(config.RootConfigActiveKey)
	if activeWorkspace != "" {
		c.active = viper.New()
		c.active.SetConfigFile(
			path.Join(
				path.Dir(c.root.ConfigFileUsed()),
				c.root.GetStringMapString(
					config.RootConfigWorkspacesKey,
				)[activeWorkspace],
			),
		)
		if err := c.active.ReadInConfig(); err != nil {
			return c.Error(err)
		}
	}

	return nil
}

type Options struct {
	viper *viper.Viper
}

type Option func(o *Options)

func WithViper(viper *viper.Viper) Option {
	return func(o *Options) {
		if viper != nil {
			o.viper = viper
		}
	}
}

func (c *ViperConfig) Write(
	o ...Option,
) error {
	opts := &Options{}
	for _, opt := range o {
		opt(opts)
	}

	v := viper.GetViper()
	if opts.viper != nil {
		v = opts.viper
	}

	if err := v.WriteConfig(); err != nil {
		return c.Error(
			err,
			WithViper(v),
		)
	}

	return nil
}
