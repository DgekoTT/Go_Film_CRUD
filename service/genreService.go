package service

import (
	"go_crud/initializers"
	"go_crud/models"
	"log"
	"strings"
)

func GetGenreIdsByName(name string) ([]*models.Genre, error) {
	namesGenre := strings.Split(name, ",")
	var genreIds []*models.Genre
	err := initializers.DB.Where("genre_name IN ?",
		namesGenre).Find(&genreIds).Error
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса:", err)
	}
	return genreIds, err
}
