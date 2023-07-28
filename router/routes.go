package router

import (
	"go_crud/controllers"
	"go_crud/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter регистрирует маршруты и обработчики в роутере Gin
func SetupRouter(r *gin.Engine) *gin.Engine {
	r.POST("/film", middleware.RequireAuth, controllers.FilmCreate)
	r.PUT("/film/id/:id", middleware.RequireAuth, controllers.FilmUpDate)
	r.GET("/films", controllers.FilmGetAll)
	r.GET("/film/id/:id", controllers.GetFilmById)
	r.DELETE("/film/id/:id", middleware.RequireAuth, controllers.FilmDelete)
	return r
}
