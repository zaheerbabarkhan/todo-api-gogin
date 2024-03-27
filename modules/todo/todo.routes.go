package todo

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine) {
	userRouter := r.Group("/todos")

	userRouter.POST("", CreateTodoHandler)
}
