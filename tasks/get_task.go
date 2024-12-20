package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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


func GetTasksList() error {
	intructions := "(echo -e 'Item\tCreated_At\tID'; jq -r '.[] | [.task.item, .task.created_at, .task.id] | @tsv' task_list.json) | column -s $'\t' -t"
	cmd := exec.Command("bash", "-c", intructions)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}
