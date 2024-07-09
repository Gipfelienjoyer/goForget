package main

import (
	"fmt"
	"goForget/src/handlers"
	"goForget/src/storage"
)

func main() {

	todos := &storage.ToDos{}
	var userInput string
	var taskId int

	fmt.Println("Welcome to your GoForget.")
	fmt.Println("Your ToDo App in the Terminal written in Go")
	fmt.Println("If you need help read the documentation or type the letter h")

	for {

		fmt.Scanln(&userInput)

		switch userInput {
		case "h":
			fmt.Println("Get some help")

		case "d":
			fmt.Println("Which Task do you want to delete? (TaskId)")
			fmt.Scanln(&taskId)
			fmt.Println("Deleting Task", handlers.DeleteTask(taskId))

		case "c":
			// Implement completion logic here

		case "ls":
			fmt.Scanln("Which Task do you want to view?", &taskId)
			fmt.Println(todos)

		case "add":
			var title, category, description string
			fmt.Println("Enter task title:")
			fmt.Scanln(&title)
			fmt.Println("Enter task category:")
			fmt.Scanln(&category)
			fmt.Println("Enter task description:")
			fmt.Scanln(&description)
			handlers.CreateTask(todos, title, category, description, nil)
			fmt.Println(todos)
			break

		default:
			fmt.Println("Please use a real command")
			fmt.Println("If you need help type the letter h")
		}

	}
}
