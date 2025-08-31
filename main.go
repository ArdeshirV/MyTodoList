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

func (t *TodoList) UpdateTaskStatus(taskID int, isDone bool) {
	for i, task := range t.Tasks {
		if task.ID == taskID {
			t.Tasks[i].IsDone = isDone
			fmt.Println("Task status updated successfully")
			return
		}
	}
	fmt.Println("Task not found to update!")
}

func (t *TodoList) deleteTask(taskID int) {
	for i, task := range t.Tasks {
		if task.ID == taskID {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			fmt.Println("Task deleted successfully")
			return
		}
	}
	fmt.Println("Task not found to delete!")
}
