package controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/models"
	"strings"
)

func FilmCreate(c *gin.Context) {

	var body struct {
		FilmName       string `json:"Body"`
		ProductionYear int16  `json:"ProductionYear"`
		Genres         string `json:"Genres"`
	}

	c.Bind(&body)
	var massGenres []*models.Genre

	for _, name := range strings.Split(body.Genres, ",") {
		massGenres = append(massGenres, &models.Genre{GenreName: name})
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

//func PostAll(c *gin.Context) {
//	posts := []models.Post{}
//	initializers.DB.Find(&posts)
//	c.JSON(200, gin.H{
//		"posts": posts,
//	})
//}
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
