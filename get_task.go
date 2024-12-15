package main

import (
	"encoding/json"
	"os"
)

func GetTasks() (TaskList, error) {
	
	dat, err := os.ReadFile(json_path)
	if err != nil {
		return TaskList{}, err
	}

	currentTasks := TaskList{}
	err = json.Unmarshal(dat, &currentTasks)
	if err != nil {
		return TaskList{}, err
	}

	return currentTasks, nil
}
