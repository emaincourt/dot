package sync

type Sync interface {
	Init() error
	AddRemote(name, remoteURL string) error
	Push(message string, files []string) error
	Pull() error
}
