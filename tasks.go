package main

import (
	"time"
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





