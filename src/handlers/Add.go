package Add

import (
	"fmt"
	"goForget/src/structs"
	"time"
)

var Title, Category, Description, Date, Time, DueDateConfirmation string
var DueDate *time.Time

func AddToDo(ToDoSlice *structs.ToDoSlice, ID int) int {

	fmt.Println("Whats the title of your Task?")
	fmt.Scanln(&Title)

	fmt.Println("In what category does your Task belong?")
	fmt.Scanln(&Category)

	fmt.Println("Whats the description of your Task?")
	fmt.Scanln(&Description)

	fmt.Println("Do you want to add an due Date? (y/n)")
	fmt.Scanln(&DueDateConfirmation)

	if DueDateConfirmation == "y" || DueDateConfirmation == "Y" {
		fmt.Println("Whats the due date?")
		fmt.Scanln(&Date)

		fmt.Println("Whats the due time")
		fmt.Scanln(&Time)
	}

	ToDoItem := structs.ToDoItem{
		ID:          ID,
		Title:       Title,
		Category:    Category,
		Description: Description,
		DueDate:     DueDate,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	ID++

	*ToDoSlice = append(*ToDoSlice, ToDoItem)

	return ID
}
