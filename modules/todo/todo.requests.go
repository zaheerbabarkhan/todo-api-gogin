package todo

import "time"

type CreateTodoRequest struct {
	Title       string    `json:"title" binding:"required,min=2,max=50"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
	Description string    `json:"description,omitempty"`
}
