package main

import (
	"errors"
	"os"
	"fmt"
)

type SubCommand interface {
	Init([]string) error
	Name() string
	Run() error
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("subcommand required")
	}

	cmds := []SubCommand{
		NewGetCommand(),
		NewAddCommand(),
		NewDeleteCommand(),
	}

	subCommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subCommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subCommand)
}
