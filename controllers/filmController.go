package controllers

import (
	"go_crud/service"
	"go_crud/utilits"

	"github.com/gin-gonic/gin"
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

func FilmUpDate(c *gin.Context) {
	service.FilmUpDate(c)
}

func FilmDelete(c *gin.Context) {
	//test ci
	id := c.Param("id")
	film, err := service.GetFilmById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	errDelete := service.FilmDelete(film)
	if errDelete != nil {
		c.JSON(500, gin.H{"error": errDelete.Error()})
		return
	}

	c.JSON(200, gin.H{"film": "удален успешно"})
}
