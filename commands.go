package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
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
	fmt.Println("test")
	fmt.Fprintf(os.Stderr, "%s\n", "hello world")
}
