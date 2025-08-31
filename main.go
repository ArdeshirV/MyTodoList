package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (t *TodoList) displayTasks() {
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

func (t *TodoList) updateTaskStatus(taskID int, isDone bool) {
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

func main() {
	todo := TodoList{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n==== Todo List Menu ====")
		fmt.Println("1. Add Tasks")
		fmt.Println("2. Display Tasks")
		fmt.Println("3. Update Tasks Status")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task title:")
			scanner.Scan()
			title := scanner.Text()
			todo.addTask(strings.TrimSpace(title))
		case "2":
			todo.displayTasks()
		case "3":
			fmt.Print("Enter task ID:")
			scanner.Scan()
			taskID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter task status (true/false): ")
			scanner.Scan()
			status := scanner.Text()
			todo.updateTaskStatus(taskID, strings.ToLower(status) == "true")
		case "4":
			fmt.Print("Enter task ID")
			scanner.Scan()
			taskID, _ := strconv.Atoi(scanner.Text())
			todo.deleteTask(taskID)
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
