package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func DeleteTask(id string) error {
	found := false
	taskList, err := GetTasks()
	if err != nil {
		return err
	}

	updateTaskList := TaskList{}

	for idx, task := range taskList {
		if task.Task.Id == id {
			found = true
			continue
		}
		updateTaskList = append(updateTaskList, taskList[idx])
	}

	if found == false {
		return fmt.Errorf("could not find task with id: %s\n", id)
	}
	
	dat, err := json.MarshalIndent(updateTaskList, "", "\t")

	if err != nil {
		return err
	}

	err = os.WriteFile(json_path, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}
