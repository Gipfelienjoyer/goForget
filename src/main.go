package main

import (
	"fmt"
	"goForget/src/handlers"
	"goForget/src/structs"
	"os"
)

func main() {

	// Used to temporarily store the Tasks before writing them to a file
	ToDoStorage := &structs.ToDoSlice{}

	// A var that stores the next ID we need to append to
	var NextID int

	var userInput string

	for {

		fmt.Scanln(&userInput)

		switch userInput {
		case "a":
			NextID = Add.AddToDo(ToDoStorage, NextID)

		case "x":
			os.Exit(1)

		default:
			fmt.Println("Please use a real command")
			fmt.Println("If you need help type the letter h")
		}

	}
}
