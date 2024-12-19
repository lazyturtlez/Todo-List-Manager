package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
)

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
		err := GetTasksList()
		if err != nil {
			return err
		}
		return nil
	}

	currentTasks, err := GetTasks()
	if err != nil {
		return err
	}
	dat, err := json.MarshalIndent(currentTasks, "", "\t")
	fmt.Println(string(dat))
	return nil
}

func NewGetCommand() *GetTasksCommand {
	gtc := &GetTasksCommand{
		fs: flag.NewFlagSet("get", flag.ContinueOnError),
	} 

	gtc.fs.BoolVar(&gtc.list, "list", false, "specifies list view")
	return gtc
}






 
