package controllers_test

import (
	"github.com/gin-gonic/gin"
	"go_crud/controllers"
)

func SetupRouterGenre() *gin.Engine {
	r := gin.Default()
	r.POST("/genre", controllers.GenreCreate)
	r.GET("/genre", controllers.GenreGetAll)
	r.DELETE("/genre/:id", controllers.GenreDelete)
	return r
}
