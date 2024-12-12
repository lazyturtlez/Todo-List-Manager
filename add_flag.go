package main

import (
	"errors"
	"flag"
	"fmt"
)



type AddTaskCommand struct {
	fs *flag.FlagSet
	multiple bool

}

func (atc *AddTaskCommand) Init(args []string) error {
	return atc.fs.Parse(args) 
}

func (atc *AddTaskCommand) Run() error {
	args := atc.fs.Args()
	
	if len(args) < 1 {
		return errors.New("Please provide a task to add")
	}

	if atc.multiple {
		for _, arg := range args {
			AddTask(arg)
			fmt.Printf("added: %s\n", arg)
		}
		return nil
	}

	if len(args) > 1 {
		return errors.New("to many args to process, use -multiple flag to process multiple tasks")
	}

	AddTask(args[0])
	fmt.Printf("added: %s\n", args[0])
	return nil
}

func (atc *AddTaskCommand) Name() string {
	return atc.fs.Name()
}

func NewAddCommand() *AddTaskCommand {
	tc := &AddTaskCommand {
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}

	tc.fs.BoolVar(&tc.multiple, "multiple", false, "allows user to add multiple task at once")

	return tc
}



