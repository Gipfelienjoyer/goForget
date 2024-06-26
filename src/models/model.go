package schema

import "time"

type Todo struct {
    ID              int                 `json:"id"`
    Title           string              `json:"title"`
    Categorie       string              `json:"categorie"`
    Description     string              `json:"note"`
    DueDate         *time.Time           `json:"dueDate"`
    CreatedAt       time.Time           `json:"createdAt"`
    CompletedAt     *time.Time           `json:"completedAt"`
}