package routes

import (
	"net/http"
	"task/controllers"
	"task/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Task Service"})
	})

	router.GET("/tasks", controllers.GetTask)

	protectedRoutes := router.Group("/")

	protectedRoutes.Use(middleware.AuthMiddleWare())

	protectedRoutes.POST("/task", controllers.AddTask)

}
