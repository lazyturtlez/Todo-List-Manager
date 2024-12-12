package main

import (
	"fmt"
	"os"
)


const json_path = "task_list.json"

func main()  {
	err := root(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

