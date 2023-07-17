package router

import (
	"github.com/gin-gonic/gin"
	"go_crud/controllers"
)

// SetupRouter регистрирует маршруты и обработчики в роутере Gin
func SetupRouter(r *gin.Engine) {
	r.POST("/film", controllers.FilmCreate)
	//r.PUT("/post/:id", controllers.PostUpDate)
	r.GET("/film", controllers.FilmGetAll)
	//r.GET("/post/:id", controllers.PostById)
	//r.DELETE("/post/:id", controllers.PostDelete)
}
