package controller

import (
	"auth/db"
	"auth/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	User_id int `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("secret_key")

func Register(context *gin.Context) {
	var req model.User

	err := context.BindJSON(&req)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to bing json check ur json again"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	_, err = db.DB.Exec("INSERT INTO users(username, email, password) VALUES($1, $2, $3)", req.Username, req.Email, hashedPassword)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func Login(context *gin.Context) {
	var req model.User

	err := context.BindJSON(&req)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json damn "})
		return
	}

	var user model.User

	err = db.DB.QueryRow("SELECT * FROM users WHERE email = $1", req.Email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	expiration := time.Now().Add(2 * time.Hour)

	claims := &Claims{
		User_id: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}
