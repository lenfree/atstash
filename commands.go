package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"srcd.works/go-git.v4/config"

	"github.com/codegangsta/cli"
	"github.com/lenfree/atstash/git"
	"github.com/lenfree/atstash/stash"
)

// GlobalFlags to expose all global flags if it exists.
var GlobalFlags = []cli.Flag{}

// Commands is a slice of cli.Command to load.
var Commands = []cli.Command{
	{
		Name:    "push",
		Aliases: []string{"p"},
		Usage:   "Push current branch to forked repo",
		Action:  cmdPush,
		Flags:   []cli.Flag{},
	},
}

// CommandNotFound returns help message.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}

func cmdPush(c *cli.Context) {

	refSpec, err := gitPush()
	if err != nil && strings.ToLower(err.Error()) != "already up-to-date" {
		log.Fatalf("Error %s pushing to remote\n", err.Error())
	}

	res := createPR(refSpec)
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

func gitPush() (config.RefSpec, error) {
	r := gitClient.New(originRepoName, forkedRepoName)
	repo, _ := r.Repo()

	var remote gitClient.Remotes

	remote = r.GetRemote(repo)
	refSpec, err := remote.PushCommit(repo)

	if err != nil {
		return refSpec, err
	}

	return refSpec, nil
}

func createPR(ref config.RefSpec) *http.Response {
	stashConfig := stashClient.StashConfig{
		User:        stashUser,
		Password:    stashPass,
		Host:        stashURL,
		ProjectKey:  projectKey,
		RepoKey:     repoKey,
		PrReviewers: reviewers,
	}

	client := stashClient.New(stashConfig)
	res, err := client.CreatePR(ref)
	if err != nil {
		fmt.Printf("Error %ss", err.Error())
	}
	return res
}
