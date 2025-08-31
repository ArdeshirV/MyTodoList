package main

import (
	"fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
	fmt.Println(Prompt("My TODO List"))
}

type Task struct {
    ID int
    Title string
    IsDone bool
}

type TodoList struct {
    Tasks []Task
}

func (t *TodoList) addTask(title string) {
    task := Task {
        ID: len(t.Tasks) + 1,
        Title: title,
        IsDone: false,
    }
    t.Tasks := t.Tasks.append(t.Tasks, task)
    fmt.Println("Task added successfully")
}


