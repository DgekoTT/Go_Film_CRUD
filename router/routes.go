package router

import (
	"github.com/gin-gonic/gin"
	"go_crud/controllers"
)

// SetupRouter регистрирует маршруты и обработчики в роутере Gin
func SetupRouter(r *gin.Engine) {
	r.POST("/film", controllers.FilmCreate)
	//r.PUT("/film/id/:id", controllers.FilmUpDate)
	r.GET("/film", controllers.FilmGetAll)
	r.GET("/film/id/:id", controllers.GetFilmById)
	r.DELETE("/film/id/:id", controllers.FilmDelete)
}
