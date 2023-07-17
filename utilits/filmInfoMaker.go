package utilits

import (
	"go_crud/models"
	"strings"
	"time"
)

type FilmInfo struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FilmName       string
	ProductionYear int16
	GenreNames     string
}

func MakeFilmInfo(films []models.Film) []FilmInfo {
	var filmInfos []FilmInfo
	for _, film := range films {
		genreNames := MakeGenres(film)

		filmInfos = append(filmInfos, CreateFilmInfo(film, genreNames))
	}
	return filmInfos
}

func MakeGenres(film models.Film) string {
	genreNames := strings.Builder{}
	if len(film.Genres) == 1 {
		return film.Genres[0].GenreName
	}
	for i, genre := range film.Genres {
		if i == len(film.Genres)-1 {
			genreNames.WriteString(genre.GenreName)
		} else {
			genreNames.WriteString(genre.GenreName)
			genreNames.WriteString(", ")
		}
	}
	return genreNames.String()
}

func CreateFilmInfo(film models.Film, genreNames string) FilmInfo {
	return FilmInfo{
		ID:             film.ID,
		CreatedAt:      film.CreatedAt,
		UpdatedAt:      film.UpdatedAt,
		FilmName:       film.FilmName,
		ProductionYear: film.ProductionYear,
		GenreNames:     genreNames,
	}
}
