package database

import (
	"fmt"

	"github.com/zaheerbabarkhan/todo-api-gogin/models"
	"gorm.io/gorm"
)

func MigrateModels(Db *gorm.DB) {
	err := Db.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Unable to migrate user model")
	}
}
