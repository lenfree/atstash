package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/lenfree/atstash/git"
	"github.com/lenfree/atstash/stash"
)

// GlobalFlags to expose all global flags if it exists.
var GlobalFlags = []cli.Flag{}

// Commands is a slice of cli.Command to load.
var Commands = []cli.Command{
	{
		Name:    "pr",
		Aliases: []string{"pr"},
		Usage:   "Create a Stash PR from remote fork",
		Action:  cmdPr,
		Flags:   []cli.Flag{},
	},
}

// CommandNotFound returns help message.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

func cmdPr(c *cli.Context) {

	branch, err := gitQuery()
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	res, err := createPR(branch)
	if err != nil {
		fmt.Printf("Error %ss", err.Error())
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("API Post Method Response Status Code: %s\n", string(res.StatusCode))
		log.Fatalf("API Post Method Response Body: %s\n", string(resBody))
	}

	if res.StatusCode == 409 {
		body := stashClient.PrStatusExists{}
		json.Unmarshal(resBody, &body)
		log.Fatalf("Error: %s", body.Errors[0].Message)
	}

	body := stashClient.PrRes{}
	json.Unmarshal(resBody, &body)
	log.Printf("PR Title: %s\n", body.Title)
	log.Printf("PR URL: %s/%s\n", stashURL, body.Link.URL)
}

func gitQuery() (string, error) {
	r := gitClient.New(originRepoName, forkedRepoName)
	repo, err := r.Repo()

	if err != nil {
		log.Fatalf("Fetch repo error: %sn", err.Error())
	}

	ref, err := r.GetHead(repo)

	branch, err := r.GetBranch(ref)

	if err != nil {
		return "", err
	}

	return branch, nil
}

func createPR(branch string) (*http.Response, error) {
	stashConfig := stashClient.StashConfig{
		User:        stashUser,
		Password:    stashPass,
		Host:        stashURL,
		ProjectKey:  projectKey,
		RepoKey:     repoKey,
		PrReviewers: reviewers,
	}

	client := stashClient.New(stashConfig)
	res, err := client.CreatePR(branch)
	if err != nil {
		return nil, err
	}
	return res, nil
}
