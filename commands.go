package main

import (
	"fmt"
	"log"
	"os"

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
	if err != nil {
		log.Fatalf("Error %s pushing to remote\n", err.Error())
	}

	// test this out
	createPR(refSpec)
}

func gitPush() (config.RefSpec, error) {
	r := gitClient.New(originRepoName, forkedRepoName)
	repo, _ := r.Repo()

	var remote gitClient.Remotes

	remote = r.GetRemote(repo)
	refSpec, err := remote.PushCommit(repo)

	if err != nil {
		return "", err
	}

	return refSpec, nil
}

func createPR(ref config.RefSpec) {
	stashClient := stashClient.New(stashUser, stashPass, stashURL, originSlug, forkedSlug, "TOP")
	resp, _ := stashClient.CreatePR("TOP", ref)
	fmt.Printf("%+#v\n", resp)
}
