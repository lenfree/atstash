package stashClient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type StashConfig struct {
	User       string
	Password   string
	Host       string
	ToSlug     string
	FromSlug   string
	FromRef    string
	ProjectKey string
}

const apiURI = "/rest/api/1.0/projects"

func New(u, p, URL, toSlug, fromSlug, projectKey string) StashConfig {
	return StashConfig{
		User:       u,
		Password:   p,
		Host:       URL,
		ToSlug:     toSlug,
		FromSlug:   fromSlug,
		ProjectKey: projectKey,
	}
}

func (s *StashConfig) Get(uri string) (*http.Response, error) {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req, err := http.NewRequest("GET", s.Host+apiURI+uri, nil)
	req.SetBasicAuth(s.User, s.Password)
	client := &http.Client{Transport: transCfg}

	resp, err := client.Do(req)

	fmt.Printf("resp %+#v", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StashConfig) CreatePR(projectKey, branch string) (*http.Response, error) {
	uri := apiURI + "/" + projectKey + "/repos/" + s.ToSlug + "/pull-requests"
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	currentBranch, err := getBranch()
	if err != nil {
		log.Fatalf("Error %s\n", err.Error())
	}

	data := Pr{
		Title:       "Talking Nerdy",
		Description: "It’s a kludge, but put the tuple from the database in the cache.",
		State:       "OPEN",
		Open:        true,
		Closed:      false,
		FromRef: Ref{
			ID: "refs/heads/" + currentBranch,
			Repository: Repo{
				Slug: s.FromSlug,
				Name: "",
				Project: Proj{
					Key: s.ProjectKey,
				},
			},
		},
		ToRef: Ref{
			ID: "refs/heads/master",
			Repository: Repo{
				Slug: "s.ToSlug",
				Name: "",
				Project: Proj{
					Key: s.ProjectKey,
				},
			},
		},
		Locked: false,
		Reviewers: []Reviewer{
			Reviewer{
				User: Usr{
					Name: "d713020",
				},
			},
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)
	req, err := http.NewRequest("POST", s.Host+apiURI+uri, body)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(s.User, s.Password)
	client := &http.Client{Transport: transCfg}

	resp, err := client.Do(req)

	fmt.Printf("resp %+#v", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getBranch() (string, error) {
	output, err := exec.Command("sh", "-c", "git branch").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
