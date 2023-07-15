package models

type Film struct {
	FilmId         int    `gorm:"primaryKey"`
	FilmName       string `gorm:"size:64"`
	ProductionYear int8
	Genres         []*Genre `gorm:"many2many:FilmGenre"`
}

type Genre struct {
	GenreId   int     `gorm:"primaryKey"`
	GenreName string  `gorm:"size:16"`
	Films     []*Film `gorm:"many2many:FilmGenre"`
}
