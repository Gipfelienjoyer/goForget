package handlers

import (
	"goForget/src/storage"
	"time"
)

func DeleteTask(taskId int) int {
	// Add deletion logic here
	return taskId
}

func CreateTask(todos *storage.ToDos, title string, category string, description string, dueDate *time.Time) {
	todos.AddToDo(title, category, description, dueDate)
}
