package main

import (
	"auth/db"
	"auth/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Auth service...")
	fmt.Println("Running on port 8000...")
	db.Connect()
	router := gin.Default()
	routes.AuthRoures(router)
	router.Run(":8000")
}
