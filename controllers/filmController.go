package controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/service"
	"go_crud/utilits"
)

func FilmCreate(c *gin.Context) {
	film, err := service.FilmCreate(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{
			"film": film,
		})
	}

}

func FilmGetAll(c *gin.Context) {
	films, err := service.FilmGetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	filmInfo := utilits.MakeFilmInfo(films)
	c.JSON(200, gin.H{
		"films": filmInfo,
	})
}

func GetFilmById(c *gin.Context) {
	id := c.Param("id")
	film, err := service.GetFilmById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"film": film,
	})
}

//func FilmUpDate(c *gin.Context) {
//	id := c.Param("id")
//	var body struct {
//		FilmName       string `json:"Body"`
//		ProductionYear int16  `json:"ProductionYear"`
//	}
//
//	errorData := c.Bind(&body)
//	if errorData != nil {
//		c.JSON(500, gin.H{"Ошибка": "Не удалось получить данные "})
//		return
//	}
//
//	var film models.Film
//	film = GetFilmById(c, id)
//
//	// Обновляем только необходимые поля фильма
//	film.FilmName = body.FilmName
//	film.ProductionYear = body.ProductionYear
//
//	errUpdate := initializers.DB.Save(&film)
//	if errUpdate.Error != nil {
//		c.JSON(500, gin.H{"Ошибка": "Не удалось обновить фильм"})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"film": film,
//	})
//}

func FilmDelete(c *gin.Context) {
	id := c.Param("id")
	film, err := service.GetFilmById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"film": "удален успешно"})
}
