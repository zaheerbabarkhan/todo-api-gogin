package user

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine) {
	userRouter := r.Group("/users")

	userRouter.POST("", CreateUserHanlder)
	userRouter.POST("/login", LoginHandler)
}
