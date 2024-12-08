package main

import (
	"encoding/json"
	"os"
	"time"
)

const json_path = "task_list.json"

func main()  {

	tasks := TaskList{
		{
			Task: TaskDetails{
				Item: "cleaning",
				CreatedAt: time.Now(),
			},
		},
		{
			Task: TaskDetails{
				Item: "homework",
				CreatedAt: time.Now().Add(time.Minute),
			},
		},
		{
			Task: TaskDetails{
				Item: "eat",
				CreatedAt: time.Now().Add(time.Minute*2),

			},
		},
	}

	f, err := os.OpenFile(json_path, os.O_RDWR | os.O_CREATE, 0)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	tasks_json, err := json.MarshalIndent(tasks, "", "\t")

	if err != nil {
		panic(err)
	}

	_, err = f.Write(tasks_json)
	if err != nil {
		panic(err)
	}
			
}

