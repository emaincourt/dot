package git

import (
	"os"
	"path/filepath"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func (s *GitSync) Push(
	message string,
	files []string,
) error {
	r, err := gogit.PlainOpen(
		s.rootDir,
	)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	for _, file := range files {
		path, err := filepath.Rel(s.rootDir, file)
		if err != nil {
			return err
		}
		_, err = w.Add(path)
		if err != nil {
			return err
		}
	}

	if message == "" {
		message = DefaultCommitMessage
	}

	commit, err := w.Commit(message, &gogit.CommitOptions{
		Author: &object.Signature{
			Name:  DefaultAuthorName,
			Email: DefaultAuthorEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	if _, err := r.CommitObject(commit); err != nil {
		return err
	}

	if err := r.Push(&gogit.PushOptions{
		Progress: os.Stdout,
	}); err != nil {
		return err
	}

	return nil
}
