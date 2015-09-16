package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/seka/wiki/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "abst",
		Usage:  "abst keyword",
		Action: command.CmdAbstruct,
		Flags: []cli.Flag{
			cli.StringFlag{},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
