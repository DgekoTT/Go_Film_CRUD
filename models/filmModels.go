package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	FilmName       string `gorm:"size:64"`
	ProductionYear int16
	Genres         []*Genre `gorm:"many2many:FilmGenre"`
}
