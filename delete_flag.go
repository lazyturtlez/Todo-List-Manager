package main

import (
	"errors"
	"flag"
	"fmt"
)

type DeleteTaskCommand struct {
	fs *flag.FlagSet
	multiple bool
}

func NewDeleteCommand() *DeleteTaskCommand {
	dc := &DeleteTaskCommand{
		fs: flag.NewFlagSet("delete", flag.ContinueOnError),
	}

	dc.fs.BoolVar(&dc.multiple, "multiple", false, "allows user to enter multiple task ids")
	return dc
}

func (dtc *DeleteTaskCommand) Init(args []string) error {
	return dtc.fs.Parse(args)
}

func (dtc *DeleteTaskCommand) Name() string {
	return dtc.fs.Name()
}

func (dtc *DeleteTaskCommand) Run() error {
	args := dtc.fs.Args()
	if len(args) < 1 {
		return errors.New("Please provid task id")
	}

	if dtc.multiple {
		for _, arg := range args {
			err := DeleteTask(arg)
			if err != nil {
				return err
			}
		}
		return nil
	}
	
	if len(args) > 1 {
		return errors.New("Please use the -multiple flag to delete more then one task")
	}
	DeleteTask(args[0])
	fmt.Printf("deleted task with id: %s\n", args[0])
	return nil
}




