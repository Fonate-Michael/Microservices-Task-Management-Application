package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	User_id int `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {

		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Heade found!"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret_key"), nil
		})

		if err != nil || !token.Valid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}

		context.Set("user_id", claims.User_id)
		context.Next()

	}
}
