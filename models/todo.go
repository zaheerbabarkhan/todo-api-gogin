package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zaheerbabarkhan/todo-api-gogin/constants"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:char(36);column:id" json:"id"`
	Title       string    `gorm:"size:50;not null;column:title" json:"title"`
	Description string    `gorm:"size:1000;column:description" json:"description"`
	DueDate     time.Time `gorm:"not null;column:due_date" json:"dueDate"`
	StatusId    int8      `gorm:"type:smallint;column:status_id;not null" json:"statusId"`
	CompletedAt time.Time `gorm:"index;column:completed_at" json:"completedAt"`
	UserID      uuid.UUID `gorm:"type:char(36);not null;column:user_id" json:"userId"`
	CreatedAt   time.Time `gorm:"not null;column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null;column:updated_at" json:"updatedAt"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) error {
	todo.StatusId = int8(constants.Status.PENDING)
	todo.CreatedAt = time.Now()
	return nil
}

func (todo *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	todo.UpdatedAt = time.Now()
	return nil
}
