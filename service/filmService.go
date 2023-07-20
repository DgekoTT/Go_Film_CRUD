package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/models"
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
