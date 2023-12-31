package service

import (
	"fmt"
	"go_crud/initializers"
	"go_crud/models"

	"github.com/gin-gonic/gin"
)

func FilmCreate(c *gin.Context) (models.Film, error) {
	var body struct {
		FilmName       string `json:"FilmName"`
		ProductionYear int16  `json:"ProductionYear"`
		Genres         string `json:"Genres"`
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось получить данные при создании фильма"})
		return models.Film{}, err
	}

	massGenres, err2 := GetGenreIdsByName(body.Genres)
	if err2 != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось получить жанры"})
		return models.Film{}, err2
	}

	film := models.Film{
		FilmName:       body.FilmName,
		ProductionYear: body.ProductionYear,
		Genres:         massGenres,
	}

	result := initializers.DB.Create(&film)

	if result.Error != nil {
		c.Status(500)
		return models.Film{}, result.Error
	}

	return film, nil
}

func FilmGetAll() ([]models.Film, error) {
	var films []models.Film
	err := initializers.DB.Preload("Genres").Find(&films).Error
	if err != nil {
		return []models.Film{}, err
	}
	return films, nil
}

func GetFilmById(id string) (models.Film, error) {
	film := models.Film{}
	err := initializers.DB.Preload("Genres").First(&film, id).Error
	if err != nil {
		return models.Film{}, err
	}
	return film, nil
}

func FilmDelete(film models.Film) error {
	if err := initializers.DB.Table("film_genre").Where("film_id = ?", film.ID).Delete(nil).Error; err != nil {
		return fmt.Errorf("не удалось удалить связи с жанрами: %v", err)
	}

	if err := initializers.DB.Delete(&film).Error; err != nil {
		return fmt.Errorf("не удалось удалить фильм: %v", err)
	}

	return nil
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

	film, err := GetFilmById(id)
	if err != nil {
		c.JSON(500, gin.H{"Ошибка": "Фильм не найден"})
		return
	}

	// Обновляем только необходимые поля фильма
	film.FilmName = body.FilmName
	film.ProductionYear = body.ProductionYear

	errUpdate := initializers.DB.Save(&film)
	if errUpdate.Error != nil {
		c.JSON(500, gin.H{"Ошибка": "Не удалось обновить фильм"})
		return
	}

	c.JSON(200, gin.H{
		"film": film,
	})
}
