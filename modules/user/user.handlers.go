package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaheerbabarkhan/todo-api-gogin/database"
	"github.com/zaheerbabarkhan/todo-api-gogin/models"
	"github.com/zaheerbabarkhan/todo-api-gogin/modules/mail"
	"github.com/zaheerbabarkhan/todo-api-gogin/types"
)

func CreateUserHanlder(c *gin.Context) {
	var userData CreateUserRequest

	fmt.Println("this is here")
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	user := models.User{
		FirstName:   userData.FirstName,
		LastName:    userData.LastName,
		Email:       userData.Email,
		Password:    userData.Password,
		AccountType: types.AccountTypes.APP,
	}

	database.Db.Create(&user)
	err := mail.SendConfirmationEmail(user.Email, "123")

	if err != nil {
		fmt.Println("Error occurred during email confirmation:", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "confirmation email not sent",
			"user":    user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"user":    user,
	})
}
