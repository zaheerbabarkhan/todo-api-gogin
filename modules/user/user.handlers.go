package user

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zaheerbabarkhan/todo-api-gogin/database"
	"github.com/zaheerbabarkhan/todo-api-gogin/models"
	"github.com/zaheerbabarkhan/todo-api-gogin/types"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHanlder(c *gin.Context) {
	var userData CreateUserRequest

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingUser models.User
	database.Db.Where("email = ?", userData.Email).Find(&existingUser)
	if existingUser.Email == userData.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exist.",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error occurred during password hash:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User cannot be created, please contact support",
		})
		return
	}

	user := models.User{
		FirstName:   userData.FirstName,
		LastName:    userData.LastName,
		Email:       userData.Email,
		Password:    string(hashedPassword),
		AccountType: types.AccountTypes.APP,
	}

	database.Db.Create(&user)

	emailIssue := os.Getenv("SMTP_ISSUE")

	if emailIssue != "" {
		// err = mail.SendConfirmationEmail(user.Email, "123")
		c.JSON(http.StatusOK, gin.H{
			"message": "user created, confirmation email not sent, please contact support.",
			"user":    user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created, confirmation email sent",
		"user":    user,
	})
}
