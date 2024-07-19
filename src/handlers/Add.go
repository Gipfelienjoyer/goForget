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

		Date := util.ScanLine("What's the due Date?", true)
		Time := util.ScanLine("What's the due Time?", true)

		TempTime, err := time.Parse("01.01.2001 15.45", Date+" "+Time)
		if err != nil {
			fmt.Println("There was an error while inputing your time. Error: " + err.Error())
		}
		fmt.Println(TempTime)
	}

	{

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
