package gitClient

import (
	"os"

	"gopkg.in/libgit2/git2go.v25"
)

type Config struct {
	Origin string
	Forked string
}

// New returns git config
func New(o, f string) Config {
	return Config{
		Origin: o,
		Forked: f,
	}
}

func (c *Config) Repo() (*git.Repository, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	repository, err := git.OpenRepository(cwd)
	if err != nil {
		return nil, err
	}
	return repository, nil

}

func (c *Config) GetRemotes(r *git.Repository) *git.RemoteCollection {
	remotes := r.Remotes
	return &remotes
}

func (c *Config) GetHead(r *git.Repository) (*git.Reference, error) {
	head, err := r.Head()
	if err != nil {
		return nil, err
	}
	return head, nil
}

func (c *Config) GetBranch(r *git.Reference) (string, error) {
	branch := r.Branch()

	name, err := branch.Name()
	if err != nil {
		return "", err
	}

	return name, nil
}
