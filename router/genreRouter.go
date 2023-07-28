package router

import (
	"go_crud/controllers"
	"go_crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouterGenre(r *gin.Engine) *gin.Engine {
	r.POST("/genre", middleware.RequireAuth, controllers.GenreCreate)
	r.GET("/genre", controllers.GenreGetAll)
	r.DELETE("/genre/:id", middleware.RequireAuth, controllers.GenreDelete)
	return r
}
