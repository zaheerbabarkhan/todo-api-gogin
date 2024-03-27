package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaheerbabarkhan/todo-api-gogin/models"
)

func CreateTodoHandler(c *gin.Context) {
	var input CreateTodoRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}
	userData, ok := c.Get("user")

	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, ok := userData.(models.User)
	if !ok {
		fmt.Println("Failed to convert user data to User struct")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
}
