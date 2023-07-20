package router

import (
	"go_crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouterGenre(r *gin.Engine) *gin.Engine {
	r.POST("/genre", controllers.GenreCreate)
	r.GET("/genre", controllers.GenreGetAll)
	r.DELETE("/genre/:id", controllers.GenreDelete)
	return r
}
