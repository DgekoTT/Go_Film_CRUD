package controllers

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/models"
	"log"
	"strings"
)

func GenreCreate(c *gin.Context) {
	var body struct {
		GenreName string `json:"genre_name"`
	}

	errBody := c.Bind(&body)
	if errBody != nil {
		c.JSON(400, gin.H{"error": errBody.Error()})
		return
	}

	genre := models.Genre{GenreName: body.GenreName}
	result := initializers.DB.Create(&genre)
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"genre": result,
	})
}

func GetGenreIdsByName(name string) []*models.Genre {
	namesGenre := strings.Split(name, ",")
	var genreIds []*models.Genre
	err := initializers.DB.Where("genre_name IN ?",
		namesGenre).Find(&genreIds).Error
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса:", err)
	}
	return genreIds
}

func GenreGetAll(c *gin.Context) {
	var genres []*models.Genre
	err := initializers.DB.Find(&genres).Error
	if err != nil {
		c.Status(500)
		return
	}

	c.JSON(500, gin.H{
		"genres": genres,
	})
}

func GenreDelete(c *gin.Context) {
	id := c.Param("id")
	err := initializers.DB.Delete(&models.Genre{}, id)
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось удалить жанр"})
		return
	}
	c.JSON(200, gin.H{
		"genre": "удален успешно",
	})
}
