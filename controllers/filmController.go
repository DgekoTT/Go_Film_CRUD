package controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/models"
	"go_crud/utilits"
)

func FilmCreate(c *gin.Context) {

	var body struct {
		FilmName       string `json:"Body"`
		ProductionYear int16  `json:"ProductionYear"`
		Genres         string `json:"Genres"`
	}

	c.Bind(&body)
	massGenres := GetGenreIdsByName(body.Genres)
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
	films := []models.Film{}
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

//
//func PostById(c *gin.Context) {
//	// get id from url
//	id := c.Param("id")
//	post := models.Post{}
//	initializers.DB.First(&post, id)
//	c.JSON(200, gin.H{
//		"post": post,
//	})
//}
//
//func PostUpDate(c *gin.Context) {
//	// get id from url
//	id := c.Param("id")
//	var body struct {
//		Body  string `json:"body"`
//		Title string `json:"title"`
//	}
//
//	c.Bind(&body)
//	var post models.Post
//	initializers.DB.First(&post, id)
//	initializers.DB.Model(&post).Updates(models.Post{
//		Title: body.Title,
//		Body:  body.Body,
//	})
//
//	c.JSON(200, gin.H{
//		"post": post,
//	})
//}
//
//func PostDelete(c *gin.Context) {
//	// get id from url
//	id := c.Param("id")
//	initializers.DB.Delete(&models.Post{}, id)
//	c.JSON(200, gin.H{
//		"post": "okey",
//	})
//}
