package git

import (
	gogit "github.com/go-git/go-git/v5"
)

func (s *GitSync) Pull() error {
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

	if err := w.Pull(&gogit.PullOptions{}); err != nil {
		if err == gogit.NoErrAlreadyUpToDate {
			return nil
		}
		return err
	}

	return nil
}
