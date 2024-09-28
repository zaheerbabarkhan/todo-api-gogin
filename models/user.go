package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/zaheerbabarkhan/todo-api-gogin/constants"
	"github.com/zaheerbabarkhan/todo-api-gogin/types"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);column:id" json:"id"`
	FirstName   string    `gorm:"size:20;not null;coulmn:first_name" json:"firstName"`
	LastName    string    `gorm:"size:20;not null;coulmn:last_name" json:"lastName"`
	Email       string    `gorm:"size:250;not null;column:email;uniqueIndex" json:"email"`
	AccountType string    `gorm:"size:10;not null;column:account_type" json:"accountType"`
	Password    string    `gorm:"coulmn:password" json:"password"`
	StatusId    int8      `gorm:"type:smallint;column:status_id;not null" json:"statusId"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()

	if user.StatusId == 0 && user.AccountType == types.AccountTypes.APP {
		user.StatusId = int8(constants.Status.PENDING)
	}

	if user.StatusId == 0 && user.AccountType == types.AccountTypes.SOCIAL {
		user.StatusId = int8(constants.Status.ACTIVE)
	}
	user.CreatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return nil
}

func (User) DefaultScope(db *gorm.DB) *gorm.DB {
	return db.Where("status_id != ?", constants.Status.DELETED)
}
