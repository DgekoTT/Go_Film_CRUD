package router

import (
	"go_crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter регистрирует маршруты и обработчики в роутере Gin
func SetupRouter(r *gin.Engine) *gin.Engine {
	r.POST("/film", controllers.FilmCreate)
	r.PUT("/film/id/:id", controllers.FilmUpDate)
	r.GET("/films", controllers.FilmGetAll)
	r.GET("/film/id/:id", controllers.GetFilmById)
	r.DELETE("/film/id/:id", controllers.FilmDelete)
	return r
}
