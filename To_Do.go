package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Description string
	Completed   bool
}

type ToDoList struct {
	Tasks []Task
}

func (tdl *ToDoList) AddTask(description string) {
	task := Task{Description: description, Completed: false}
	tdl.Tasks = append(tdl.Tasks, task)
}

func (tdl *ToDoList) MarkTaskCompleted(index int) {
	if index >= 0 && index < len(tdl.Tasks) {
		tdl.Tasks[index].Completed = true
	} else {
		fmt.Println("Invalid Task Index")
	}
}

func (tdl *ToDoList) PrintTasks() {
	fmt.Println("To-Do List:")
	for i, task := range tdl.Tasks {
		status := "[ ]"
		if task.Completed {
			status = "[✔]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
	}
}

func main() {
	todoList := ToDoList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a command (add/mark/print/exit): ")
		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "add":
			fmt.Print("Enter Task Description: ")
			scanner.Scan()
			description := scanner.Text()
			todoList.AddTask(description)
		case "mark":
			fmt.Print("Enter Task Index to Mark as Completed: ")
			scanner.Scan()
			index, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid Task Index")
			} else {
				todoList.MarkTaskCompleted(index - 1)
			}
		case "print":
			todoList.PrintTasks()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Command")
		}
	}
}