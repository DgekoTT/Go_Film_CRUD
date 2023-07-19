package controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/models"
	"go_crud/utilits"
)

func FilmCreate(c *gin.Context) {

	var body struct {
		FilmName       string `json:"FilmName"`
		ProductionYear int16  `json:"ProductionYear"`
		Genres         string `json:"Genres"`
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось получить данные при создании фильма"})
		return
	}

	massGenres := GetGenreIdsByName(body.Genres)
	if massGenres == nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось получить жанры"})
		return
	}
	film := models.Film{FilmName: body.FilmName, ProductionYear: body.ProductionYear,
		Genres: massGenres}

	result := initializers.DB.Create(&film)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"film": film,
	})
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
	initializers.DB.First(&film, id)
	c.JSON(200, gin.H{
		"post": film,
	})
}

func FilmUpDate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		FilmName       string `json:"Body"`
		ProductionYear int16  `json:"ProductionYear"`
	}

	errorData := c.Bind(&body)
	if errorData != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось получить данные "})
		return
	}

	var film models.Film
	err := initializers.DB.First(&film, id)
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалоось найти фильм"})
		return
	}
	errUpdate := initializers.DB.Model(&film).Updates(models.Film{
		FilmName:       body.FilmName,
		ProductionYear: body.ProductionYear,
	})
	if errUpdate != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось обновить фильм"})
		return
	}

	c.JSON(200, gin.H{
		"film": film,
	})
}

func FilmDelete(c *gin.Context) {
	id := c.Param("id")
	err := initializers.DB.Delete(&models.Film{}, id)
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось удалить фильм"})
		return
	}
	c.JSON(200, gin.H{
		"film": "удален успешно",
	})
}
