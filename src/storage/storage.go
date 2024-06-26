package storage

import (
	"goForget/src/models"
	"golang.org/x/text/date"
	"time"
)

// Defining a slice of ToDo-Structs which are definded in src/schema
type ToDos []schema.Todo

// An INT that specifies where the next toDo item needs to be placed
var nextID int

func (t *ToDos) addToDo(title string, categorie string, description string, dueDate *time.Time) {
	toDoItem := schema.Todo {
		ID: nextID,
		Title: title,
		Categorie: categorie,
		Description: description,
		DueDate: dueDate,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	nextID ++

	*t = append(*t, toDoItem)
}