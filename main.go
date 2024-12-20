package main

import (
	"fmt"
	"os"
	
	"github.com/lazyturtlez/todoCli/flags"
)

func main()  {
	err := flags.Root(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

