package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zaheerbabarkhan/todo-api-gogin/database"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		os.Exit(1)
	}

	database.ConnectDatabase()

	router.GET("status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
