package todo

import "time"

type CreateTodoRequest struct {
	Title       string    `form:"title" binding:"required,min=2,max=50"`
	DueDate     time.Time `form:"dueDate" binding:"required"`
	Description string    `form:"description,omitempty"`
}
