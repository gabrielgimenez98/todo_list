package main

import "fmt"

type task struct {
	title  string
	detail string
	status int
}

var task_list []task

func main() {
	fmt.Println("hello")
}

func add_task(new_task task) {
	task_list = append(task_list, new_task)
}
