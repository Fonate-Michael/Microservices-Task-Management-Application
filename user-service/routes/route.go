package routes

import (
	"auth/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoures(r *gin.Engine) {

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/health", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Bitch am alive!"})
	})

}
