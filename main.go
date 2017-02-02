package main

import (
	"os"

	"github.com/codegangsta/cli"
)

type config struct {
	StashURL     string   `yaml:"stashurl"`
	RemoteOrigin string   `yaml:"remote_orgin"`
	RemoteFork   string   `yaml:"remote_fork"`
	Users        []string `yaml:"users"`
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
