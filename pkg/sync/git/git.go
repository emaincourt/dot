package git

const (
	GitDirPath = ".git"
)

const (
	DefaultCommitMessage = "[Dot] Update"
	DefaultAuthorName    = "Dot"
	DefaultAuthorEmail   = "dot@dot.com"
)

type GitSync struct {
	rootDir string
}

func NewGitSync(rootDir string) *GitSync {
	return &GitSync{
		rootDir: rootDir,
	}
}
