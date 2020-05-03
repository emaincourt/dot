package viper

import (
	"github.com/emaincourt/dot/pkg/config"
)

func (c *ViperConfig) SetEnvironmentVariableForActiveWorkspace(
	name,
	value string,
	private bool,
) error {
	var envs []config.WorkspaceEnv
	if err := c.active.UnmarshalKey(config.ActiveWorkspaceEnvKey, &envs); err != nil {
		return c.Error(err)
	}

	for index, env := range envs {
		if env.Name == name {
			envs = append(envs[:index], envs[index+1:]...)
		}
	}

	envs = append(envs, config.WorkspaceEnv{
		Name:    name,
		Value:   value,
		Private: private,
	})
	c.active.Set(config.ActiveWorkspaceEnvKey, envs)

	return c.Write(WithViper(c.active))
}

func (c *ViperConfig) UnSetEnvironmentVariableForActiveWorkspace(
	name string,
) error {
	var envs []config.WorkspaceEnv
	if err := c.active.UnmarshalKey(
		config.ActiveWorkspaceEnvKey,
		&envs,
	); err != nil {
		return c.Error(err)
	}

	var newEnvs []config.WorkspaceEnv
	for _, env := range envs {
		if env.Name != name {
			newEnvs = append(newEnvs, env)
		}
	}
	c.active.Set(config.ActiveWorkspaceEnvKey, newEnvs)

	return c.Write(WithViper(c.active))
}
