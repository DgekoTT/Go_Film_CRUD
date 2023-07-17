package models

type Genre struct {
	GenreId   int     `gorm:"primaryKey;autoIncrement"`
	GenreName string  `gorm:"size:16"`
	Films     []*Film `gorm:"many2many:FilmGenre"`
}
