package tasks

import (
	"encoding/json"
	"time"
	"os"
	"github.com/google/uuid"
)

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
