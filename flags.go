package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)


type SubCommand interface {
	Init([]string) error
	Name() string
	Run() error
}


type GetTasksCommand struct {
	fs *flag.FlagSet
	list bool
}

func (gtc *GetTasksCommand) Init(args []string) error {
	return gtc.fs.Parse(args)
}

func (gtc *GetTasksCommand) Name() string{
	return gtc.fs.Name()
}

func (gtc *GetTasksCommand) Run() error {

	if len(gtc.fs.Args()) != 0 {
		return errors.New("no arguments are accepted")
	}

	if gtc.list {
		println("Should be in list format")
		return nil
	}

	currentTasks, err := GetTasks(json_path)
	if err != nil {
		return err
	}
	fmt.Println(currentTasks)
	return nil
}

func NewGetCommand() *GetTasksCommand {
	gtc := &GetTasksCommand{
		fs: flag.NewFlagSet("get", flag.ContinueOnError),
	} 

	gtc.fs.BoolVar(&gtc.list, "list", false, "specifies list view")
	return gtc
}


func root(args []string) error {
	if len(args) < 1 {
		return errors.New("subcommand required")
	}

	cmds := []SubCommand{
		NewGetCommand(),
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




 
