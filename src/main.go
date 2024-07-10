package main

import (
	"fmt"
	"goForget/src/handlers"
	"goForget/src/structs"
)

func main() {

	type ToDoSlice []structs.ToDoItem
	ToDoStorage := &ToDoSlice{}
	nextId := 0

	var userInput string
	var taskId int

	for {

		fmt.Scanln(&userInput)

		switch userInput {
		case "u":
			fmt.Scanln(&taskId)
			handlers.UpdateTask(todos, taskId)

		case "h":
			fmt.Println("Get some help")

		case "d":
			fmt.Scanln(&taskId)
			handlers.DeleteTask(todos, taskId)

		case "c":
			// Implement completion logic here

		case "ls":
			fmt.Scanln("Which Task do you want to view?", &taskId)
			fmt.Println(todos)

		case "add":
			break

		default:
			fmt.Println("Please use a real command")
			fmt.Println("If you need help type the letter h")
		}

	}
}
