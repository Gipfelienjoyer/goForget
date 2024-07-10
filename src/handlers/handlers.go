package handlers

import (
	"goForget/src/storage"
	"time"
)

func DeleteTask(todos *storage.ToDos, taskId int) int {
	todo
	return taskId
}

func CreateTask(todos *storage.ToDos, title string, category string, description string, dueDate *time.Time) {
	todos.AddToDo(title, category, description, dueDate)
}
