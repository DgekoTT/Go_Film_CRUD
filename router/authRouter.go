package router

import (
	"go_crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine) *gin.Engine {
	r.POST("/auth/singup", controllers.Signup)
	r.GET("/auth", controllers.GetAllUsers)
	// r.DELETE("/genre/:id", controllers.GenreDelete)
	return r
}
