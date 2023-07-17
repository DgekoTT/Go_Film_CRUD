package models

type Film struct {
	FilmId         int    `gorm:"primaryKey;autoIncrement"`
	FilmName       string `gorm:"size:64"`
	ProductionYear int16
	Genres         []*Genre `gorm:"many2many:FilmGenre"`
}
