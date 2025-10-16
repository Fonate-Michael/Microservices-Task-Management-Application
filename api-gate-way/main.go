package main

import (
	"fmt"
	"gate/proxy"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("User Service Started with success...")

	const UserService = "http://localhost:8000"
	const TaskService = "http://localhost:8001"

	router := gin.Default()
	router.Use(cors.Default())

	router.Any("/users/*path", proxy.ReverseProxy(UserService))
	router.Any("/task/*path", proxy.ReverseProxy(TaskService))

	router.Run(":8002")

}
