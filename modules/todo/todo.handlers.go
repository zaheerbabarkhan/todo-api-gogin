package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(c *gin.Context) {
	var input CreateTodoRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
