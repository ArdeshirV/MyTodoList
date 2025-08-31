package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Prompt("My TODO List"))
}

type Task struct {
	ID     int
	Title  string
	IsDone bool
}

type TodoList struct {
	Tasks []Task
}

func (t *TodoList) addTask(title string) {
	task := Task{
		ID:     len(t.Tasks) + 1,
		Title:  title,
		IsDone: false,
	}
	t.Tasks = append(t.Tasks, task)
	fmt.Println("Task added successfully")
}

func (t *TodoList) displayTask() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks found!")
		return
	}

	fmt.Println("Tasks:")
	var doneStatus string
	for _, task := range t.Tasks {
		if task.IsDone {
			doneStatus = "X"
		} else {
			doneStatus = " "
		}
		fmt.Printf("[%s] %d. %s\n", doneStatus, task.ID, task.Title)
	}
}
