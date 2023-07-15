package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	filmId         int    `gorm:"primaryKey"`
	filmName       string `gorm:"size:64"`
	productionYear int8
	Genres         []Genre `gorm:"many2many:film_genres"`
}

type Genre struct {
	gorm.Model
	GenreId   int    `gorm:"primaryKey"`
	GenreName string `gorm:"size:16"`
	Films     []Film `gorm:"many2many:film_genres"`
}

type FilmGenre struct {
	FilmId  int `gorm:"primaryKey"`
	GenreId int
}
