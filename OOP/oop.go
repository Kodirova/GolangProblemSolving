package main

import (
	"fmt"
)

type task struct {
	title     string
	step      string
	task_step string
}

type Director interface {
	Givetask(id int, err error)
}

type TeamLead interface {
	Delegate(status string, err error)
}

type Programmer interface {
	Develop(err error)
}

func (t *task) Givetask(id int, err error) {
	fmt.Printf("Given task %s", t.title)
	if err != nil {
		fmt.Println("Error occured", err)
	}
}

func (t *task) Delegate() {
	fmt.Printf("Given step %s", t.step)
}

func (t *task) Develop() {
	fmt.Printf("Given task step %s", t.task_step)
}
