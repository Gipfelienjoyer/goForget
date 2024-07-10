package handlers

import (
	"fmt"
	"goForget/src/structs"
)

type ToDoSlice struct {
	structs.ToDoSlice
}

var Title, Category, Description, Date, Time string


func (ToDoSliceVar ToDoSlice) addToDo(ID int) {

	fmt.Println("Whats the title of your Task?")
	fmt.Scanln(&Title)

	fmt.Println("In what category does your Task belong?")
	fmt.Scanln(&Category)

	fmt.Println("Whats the description of your Task?")
	fmt.Scanln(&Description)


	ToDoItem := structs.ToDoItem{
		ID:
	}
}
