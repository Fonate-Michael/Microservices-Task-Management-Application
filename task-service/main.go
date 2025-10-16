package main

import (
	"fmt"
	"task/db"
	"task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Connected to Task Service hehehe")
	fmt.Println("Running on port 8001...")
	db.Connect()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8001")
}
