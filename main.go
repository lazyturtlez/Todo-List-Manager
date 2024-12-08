package main

import "fmt"


const json_path = "task_list.json"

func main()  {
	err := DeleteTask(json_path, "1")
	if err != nil {
		panic(err)
	}
	fmt.Println("task list updated")
}

