package Add

import (
	"goForget/src/storage"
	"time"
)

func DeleteTask(todos *storage.ToDos, taskId int) int {
	todos.DeleteToDo(taskId)
	return taskId
}

func CreateTask(todos *storage.ToDos, title string, category string, description string, dueDate *time.Time) {
	todos.AddToDo(title, category, description, dueDate)
}

func UpdateTask(todos *storage.ToDos, taskId int) {
	todos.UpdateToDo(taskId, "update", "", "", nil)
}
