package git

import (
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func (s *GitSync) AddRemote(
	name,
	remoteURL string,
) error {
	r, err := gogit.PlainOpen(
		s.rootDir,
	)
	if err != nil {
		return err
	}

	rm, err := r.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{remoteURL},
	})
	if err != nil {
		return err
	}

	rm.Fetch(&gogit.FetchOptions{
		RemoteName: name,
	})

	return nil
}
