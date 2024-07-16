package Add

import (
	"fmt"
	"goForget/src/structs"
	"goForget/src/util"
	"time"
)

var Title, Category, Description, Date, Time, DueDateConfirmation string
var DueDate *time.Time

func AddToDo(ToDoSlice *structs.ToDoSlice, ID int) int {

	Title = util.ScanLine("Whats the title of your Task", true)
	Category = util.ScanLine("In what category does your task belong?", true)
	Description = util.ScanLine("Whats the description of your task", true)
	DueDateConfirmation = util.ScanLine("Do you want to add an due Date? (y/n)", true)

	if DueDateConfirmation == "y" || DueDateConfirmation == "Y" {
		fmt.Println("Whats the due date?")
		for true {
			fmt.Scanln(&Date)

			if Date {

			}
		}
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
