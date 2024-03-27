package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/zaheerbabarkhan/todo-api-gogin/middleware"
)

func SetUpRoutes(r *gin.Engine) {
	todoRouter := r.Group("/todos")
	todoRouter.Use(middleware.AuthRequired)
	todoRouter.POST("", CreateTodoHandler)
}
