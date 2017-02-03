package main

import (
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/joho/godotenv"
)

var stashUser string
var stashPass string
var stashURL string
var originRepoName string
var forkedRepoName string
var originSlug string
var forkedSlug string
var users []string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	stashUser = os.Getenv("STASH_USERNAME")
	stashPass = os.Getenv("STASH_PASSWORD")
	stashURL = os.Getenv("STASH_URL")
	originRepoName = os.Getenv("ORIGIN_REPO_NAME")
	forkedRepoName = os.Getenv("FORKED_REPO_NAME")
	originSlug = os.Getenv("ORIGIN_REPO_SLUG")
	forkedSlug = os.Getenv("FORKED_REPO_SLUG")
	users = strings.Split(os.Getenv("USERS"), ":")
	users = strings.Split(os.Getenv("USERS"), ":")
}

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "lenfree"
	app.Email = "lenfree.yeung@gmail.com"
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
