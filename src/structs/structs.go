package structs

import "time"

type ToDoItem struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Category    string     `json:"category"`
	Description string     `json:"note"`
	DueDate     *time.Time `json:"dueDate"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
}

type ToDoSlice []ToDoItem
