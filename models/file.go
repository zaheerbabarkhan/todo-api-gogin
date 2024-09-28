package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zaheerbabarkhan/todo-api-gogin/constants"
	"gorm.io/gorm"
)

type TodoFile struct {
	ID        uuid.UUID `gorm:"type:char(36);column:id" json:"id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	TodoID    uuid.UUID `gorm:"type:char(36);not null;column:todo_id" json:"todoId"`
	StatusId  int8      `gorm:"column:status_id;not null" json:"statusId"`
	SignedUrl string    `gorm:"-" json:"signedUrl"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
	Todo      Todo      `gorm:"foreignKey:TodoID" json:"todo"`
}

func (file *TodoFile) BeforeCreate(tx *gorm.DB) error {
	file.StatusId = int8(constants.Status.PENDING)
	file.CreatedAt = time.Now()
	return nil
}

func (file *TodoFile) BeforeUpdate(tx *gorm.DB) (err error) {
	file.UpdatedAt = time.Now()
	return nil
}
