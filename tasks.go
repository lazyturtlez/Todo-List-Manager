package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
)

type TaskWrapper struct {
	Task TaskDetails `json:"task"`
}


type TaskDetails struct {
	Item string `json:"item"`
	CreatedAt time.Time `json:"created_at"`
	Id string `json:"id"`
}

type TaskList []TaskWrapper


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

func AddTask(title string) error { 

	newTask := TaskWrapper{
		Task: TaskDetails{
			Item: title,
			CreatedAt: time.Now(),
			Id: uuid.NewString(),
		},
	}

	taskList, err := GetTasks()
	if err != nil {
		return err
	}

	taskList = append(taskList, newTask)
	dat, err := json.MarshalIndent(taskList, "", "\t")

	if err != nil {
		return err
	}

	err = os.WriteFile(json_path, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {

	taskList, err := GetTasks()
	if err != nil {
		return err
	}

	updateTaskList := TaskList{}

	for idx, task := range taskList {
		if task.Task.Id == id {
			continue
		}
		updateTaskList = append(updateTaskList, taskList[idx])
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



