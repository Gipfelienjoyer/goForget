package main

import (
	"fmt"
	"goForget/src/handlers"
	"goForget/src/models"
)


func main () {
	
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
			fmt.Println("Deleting Task", handlers.CreateTask(taskId))

		case "c":
			fmt.Println("Creating")

		case "ls":
			fmt.Println("List")
		}

		
	}
}

