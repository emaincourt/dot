package config

const (
	DefaultConfigFileType      = "yaml"
	DefaultConfigFileExtension = ".yaml"

	DefaultWorkspacesDir = "workspaces"
)

const (
	RootConfigWorkspacesKey = "workspaces"
	RootConfigActiveKey     = "active"

	ActiveWorkspaceConfigSourcesKey = "sources"
	ActiveWorkspaceEnvKey           = "env"
)

type WorkspaceName = string

type RootConfig struct {
	Active     WorkspaceName            `yaml:"active"`
	Workspaces map[WorkspaceName]string `yaml:"workspaces"`
}

type WorkspaceSource = string

type WorkspaceEnv = struct {
	Name    string `yaml:"name"`
	Value   string `yaml:"value"`
	Private bool   `yaml:"private"`
}

type WorkspaceConfig struct {
	Sources []WorkspaceSource `yaml:"sources"`
	Env     []WorkspaceEnv    `yaml:"env"`
}

type Config interface {
	Init(cfgPath string) error
	Load(cfgPath string) error
	GetWorkspaces() (map[string]string, error)
	CreateWorkspace(name string) error
	DeleteWorkspace(name string) error
	GetActiveWorkspace() (*WorkspaceConfig, error)
	GetSourcesForActiveWorkspace() ([]WorkspaceSource, error)
	AddSourceForActiveWorkspace(source string) error
	SetActiveWorkspace(name string) error
	SetEnvironmentVariableForActiveWorkspace(name, value string, private bool) error
	UnSetEnvironmentVariableForActiveWorkspace(name string) error
	Tidy() error
}
