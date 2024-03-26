package user

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zaheerbabarkhan/todo-api-gogin/constants"
	"github.com/zaheerbabarkhan/todo-api-gogin/database"
	"github.com/zaheerbabarkhan/todo-api-gogin/models"
	"github.com/zaheerbabarkhan/todo-api-gogin/types"
	"github.com/zaheerbabarkhan/todo-api-gogin/utils"
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

func LoginHandler(c *gin.Context) {
	var loginData LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	database.Db.Where("email = ? AND account_type = ?", loginData.Email, types.AccountTypes.APP).Find(&user)

	if user.Email != loginData.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials.1",
		})
		return
	}

	if user.StatusId == int8(constants.Status.PENDING) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Confirm your email please.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials.2",
		})
		return
	}

	token, err := utils.IssueJWTToken(string(user.ID.String()))

	if err != nil {
		fmt.Println("Error occurred during jwt sign:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot login, please try again.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login successfull",
		"token":   token,
	})

}
