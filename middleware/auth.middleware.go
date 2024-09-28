package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zaheerbabarkhan/todo-api-gogin/database"
	"github.com/zaheerbabarkhan/todo-api-gogin/models"
	"github.com/zaheerbabarkhan/todo-api-gogin/utils"
)

func AuthRequired(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	token = strings.Split(token, " ")[1]

	jwtToken, err := jwt.ParseWithClaims(token, &utils.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !jwtToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Token",
		})
		return
	}
	if claims, ok := jwtToken.Claims.(*utils.JWTClaims); ok {
		if claims.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token Expired",
			})
			return
		}
		var user models.User
		database.Db.Where("id = ?", claims.UserId).Find(&user)

		if user.ID.String() != claims.UserId {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Token",
		})
		return
	}

}
