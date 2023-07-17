package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	GenreName string `gorm:"size:16"`
}
