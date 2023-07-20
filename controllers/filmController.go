package controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/models"
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
	var films []models.Film
	err := initializers.DB.Preload("Genres").Find(&films).Error
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось получить фильмы"})
		return
	}

	filmInfo := utilits.MakeFilmInfo(films)
	c.JSON(200, gin.H{
		"films": filmInfo,
	})
}

func GetFilmById(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	film := models.Film{}
	initializers.DB.Preload("Genres").First(&film, id)
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
	film := models.Film{}
	if err := initializers.DB.First(&film, id).Error; err != nil {
		c.JSON(404, gin.H{"Ошибка": "Фильм не найден"})
		return
	}
	if err := initializers.DB.Table("film_genre").Where("film_id = ?", film.ID).Delete(nil).Error; err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось удалить связи с жанрами"})
		return
	}
	if err := initializers.DB.Delete(&film).Error; err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось удалить фильм"})
		return
	}

	c.JSON(200, gin.H{"film": "удален успешно"})
}
