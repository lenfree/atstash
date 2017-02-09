package stashClient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type StashConfig struct {
	User        string
	Password    string
	Host        string
	FromRef     string
	ProjectKey  string
	RepoKey     string
	PrReviewers []string
}

const apiURI = "/rest/api/1.0/projects"

func New(c StashConfig) StashConfig {
	return StashConfig{
		User:        c.User,
		Password:    c.Password,
		Host:        c.Host,
		ProjectKey:  c.ProjectKey,
		RepoKey:     c.RepoKey,
		PrReviewers: c.PrReviewers,
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

func (s *StashConfig) CreatePR(branch string) (*http.Response, error) {
	uri := apiURI + "/" + s.ProjectKey + "/repos/" + s.RepoKey + "/pull-requests"

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// ref returns refs/heads/branch:refs/heads/branch
	fromRef := "refs/heads/" + branch

	fromRefUser := "~" + s.User

	var reviewers []Reviewer

	for _, u := range s.PrReviewers {
		reviewer := Reviewer{
			User: Usr{
				Name: u,
			},
		}
		reviewers = append(reviewers, reviewer)
	}

	data := Pr{
		Title:       branch,
		Description: "",
		State:       "OPEN",
		Open:        true,
		Closed:      false,
		FromRef: Ref{
			ID: fromRef,
			Repository: Repo{
				Slug: s.RepoKey,
				Name: "",
				Project: Proj{
					Key: fromRefUser,
				},
			},
		},
		ToRef: Ref{
			ID: "refs/heads/master",
			Repository: Repo{
				Slug: s.RepoKey,
				Name: "",
				Project: Proj{
					Key: s.ProjectKey,
				},
			},
		},
		Locked:    false,
		Reviewers: reviewers,
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(data)
	req, err := http.NewRequest("POST", s.Host+uri, body)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.SetBasicAuth(s.User, s.Password)
	client := &http.Client{Transport: transCfg}

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("err %+#v", err.Error())
		return nil, err
	}

	return res, nil
}
