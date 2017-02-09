package gitClient

import (
	"os"
	"strings"

	"srcd.works/go-git.v4"
	"srcd.works/go-git.v4/plumbing"
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
	repository, err := git.PlainOpen(cwd)
	if err != nil {
		return nil, err
	}
	return repository, nil

}

func (c *Config) GetHead(r *git.Repository) (*plumbing.Reference, error) {
	head, err := r.Head()
	if err != nil {
		return nil, err
	}
	return head, nil
}

func (c *Config) GetBranch(r *plumbing.Reference) string {
	ref := r.Name()

	// ref returns refs/heads/branch:refs/heads/branch
	branch := strings.SplitAfterN(strings.Split(ref.String(), ":")[0], "/", 3)

	return branch[len(branch)-1]
}
