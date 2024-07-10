package storage

import (
	"time"
)

type Item struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Category    string     `json:"category"`
	Description string     `json:"note"`
	DueDate     *time.Time `json:"dueDate"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

// Defining a slice of ToDo-Structs which are defined in src/models
type ToDos []Item

// An INT that specifies where the next todo
var nextID int

func (t *ToDos) AddToDo(title string, category string, description string, dueDate *time.Time) {
	toDoItem := Item{
		ID:          nextID,
		Title:       title,
		Category:    category,
		Description: description,
		DueDate:     dueDate,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	nextID++
	*t = append(*t, toDoItem)
}

func (t *ToDos) DeleteToDo(taskId int) {

}

func (t *ToDos) UpdateToDo(taskId int, title string, category string, description string, dueDate *time.Time) {

	list := *t

	if len(title) != 0 {
		list[taskId].Title = title
	}
}
