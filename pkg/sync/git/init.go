package git

import (
	gogit "github.com/go-git/go-git/v5"
)

func (s *GitSync) Init() error {
	_, err := gogit.PlainInit(
		s.rootDir,
		false,
	)
	if err != nil {
		return err
	}

	return nil
}
