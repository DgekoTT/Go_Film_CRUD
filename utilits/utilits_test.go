package utilits_test

import (
	"testing"

	_ "github.com/stretchr/testify/assert"

	"go_crud/models"
	"go_crud/utilits"
)

func TestMakeGenres(t *testing.T) {
	film := models.Film{
		Genres: []*models.Genre{
			{GenreName: "комедия"},
			{GenreName: "драма"},
			{GenreName: "триллер"},
		},
	}

	genres := utilits.MakeGenres(film)

	if genres != "комедия, драма, триллер" {
		t.Errorf("expected %s, got %s", "комедия, драма, триллер", genres)
	}

	film2 := models.Film{
		Genres: []*models.Genre{
			{GenreName: "комедия"},
		},
	}
	if utilits.MakeGenres(film2) != "комедия" {
		t.Errorf("expected %s, got %s", "комедия", genres)
	}

}

//func TestMakeFilmInfo(t *testing.T) {
//	films := []models.Film{
//		{
//			FilmName:       "Облака",
//			ProductionYear: 2000,
//			Genres: []*models.Genre{
//				{GenreName: "комедия"},
//				{GenreName: "драма"},
//				{GenreName: "триллер"},
//			},
//		},
//		{
//			FilmName:       "Сон",
//			ProductionYear: 2010,
//			Genres: []*models.Genre{
//				{GenreName: "комедия"},
//			},
//		},
//	}
//
//	info := utilits.MakeFilmInfo(films)
//	result := []utilits.FilmInfo{
//		{
//,			FilmName:       "Облака",
//            ProductionYear: 2000,
//            Genres:         "комедия, драма, триллер",
//        },
//		{
//            FilmName:       "Сон",
//            ProductionYear: 2010,
//            Genres:         "комедия",
//        },
//	}
//		assert.Equal(t, "Облака (2000), комедия,")
//}
