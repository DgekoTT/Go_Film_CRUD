package router

import (
	"go_crud/controllers"
	"go_crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine) *gin.Engine {
	r.POST("/auth/signup", controllers.Signup)
	r.GET("/auth", middleware.RequireAuth, controllers.GetAllUsers)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/logout", controllers.Logout)
	return r
}
