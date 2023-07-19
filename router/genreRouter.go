package router

import (
	"github.com/gin-gonic/gin"
	"go_crud/controllers"
)

func SetupRouterGenre(r *gin.Engine) {
	r.POST("/genre", controllers.GenreCreate)
	r.GET("/genre", controllers.GenreGetAll)
	r.DELETE("/genre/:id", controllers.GenreDelete)
}
