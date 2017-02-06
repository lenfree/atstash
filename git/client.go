package gitClient

import (
	"log"
	"os"

	"strings"

	git "srcd.works/go-git.v4"
	"srcd.works/go-git.v4/config"
)

type Config struct {
	Origin string
	Forked string
}

type Remotes struct {
	Origin    string
	OriginURL string
	Remote    string
	RemoteURL string
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
		log.Fatalf("Error %s\n", err.Error())
	}
	repository, err := git.PlainOpen(cwd)
	if err != nil {
		return nil, err
	}
	return repository, nil

}

// GetRemote returns origin and forked remote
func (c *Config) GetRemote(r *git.Repository) Remotes {

	var origin string
	var remote string
	var originURL string
	var remoteURL string

	remotes, _ := r.Remotes()

	for _, r := range remotes {
		rConfig := r.Config()

		if toLower(c.Forked) == toLower(rConfig.Name) {
			remote = rConfig.Name
			remoteURL = rConfig.URL
		}

		if toLower(c.Origin) == toLower(rConfig.Name) {
			origin = rConfig.Name
			originURL = rConfig.URL
		}
	}
	return Remotes{
		Origin:    origin,
		OriginURL: originURL,
		Remote:    remote,
		RemoteURL: remoteURL,
	}
}

func (c *Remotes) PushCommit(r *git.Repository) (config.RefSpec, error) {

	headRef, _ := r.Head()
	refSpec := config.RefSpec(headRef.Name() + ":" + headRef.Name())

	err := r.Push(&git.PushOptions{
		RemoteName: c.Remote,
		RefSpecs:   []config.RefSpec{refSpec},
	})
	if err != nil {
		return refSpec, err
	}
	return refSpec, nil
}

func toLower(s string) string {
	return strings.ToLower(s)
}
