package router

import (
	"github.com/gin-gonic/gin"
	"go_crud/controllers"
)

func SetupRouterGenre(r *gin.Engine) {
	r.POST("/genre", controllers.GenreCreate)
	r.GET("/genre", controllers.GenreGetAll)
	//r.PUT("/post/:id", controllers.PostUpDate)
	//r.GET("/posts", controllers.PostAll)
	//r.GET("/post/:id", controllers.PostById)
	//r.DELETE("/post/:id", controllers.PostDelete)
}
