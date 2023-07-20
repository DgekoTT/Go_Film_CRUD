package controllers_test

import (
	"bytes"
	"encoding/json"
	"go_crud/controllers"
	"go_crud/router"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"testing"
)

func createGenre(t *testing.T, router *gin.Engine, genreName string) *httptest.ResponseRecorder {
	genreBody := map[string]string{"genre_name": genreName}
	bodyBytes, err := json.Marshal(genreBody)
	assert.NoError(t, err)

	genreReq, err := http.NewRequest("POST", "/genre", bytes.NewBuffer(bodyBytes))
	assert.NoError(t, err)
	genreReq.Header.Set("Content-Type", "application/json")

	genreRes := httptest.NewRecorder()
	router.ServeHTTP(genreRes, genreReq)

	return genreRes
}

// Проверка создания жанра с уникальным именем
func TestCreateGenre(t *testing.T) {
	r := gin.Default()
	router := router.SetupRouterGenre(r)

	genres := []string{"комедия", "драма"}
	for _, genre := range genres {
		genreRes := createGenre(t, router, genre)
		assert.Equal(t, http.StatusOK, genreRes.Code)
	}

	// Попытка создания жанра с уже существующим именем
	failGenre := createGenre(t, router, "комедия")
	assert.Equal(t, http.StatusBadRequest, failGenre.Code)

}

func TestGetGenreIdsByName(t *testing.T) {
	name := "комедия"
	genreIds := controllers.GetGenreIdsByName(name)
	expectId := uint(1)
	assert.Equal(t, expectId, genreIds[0].ID)

	name = "драма"
	genreIds = controllers.GetGenreIdsByName(name)
	assert.NotEqual(t, expectId, genreIds[0].ID)
}

func TestGenreGetAll(t *testing.T) {
	r := gin.Default()
	router := router.SetupRouterGenre(r)

	genreReq, err := http.NewRequest("GET", "/genre", nil)
	assert.NoError(t, err)
	genreReq.Header.Set("Content-Type", "application/json")
	genreRes := httptest.NewRecorder()
	router.ServeHTTP(genreRes, genreReq)
	assert.Equal(t, http.StatusOK, genreRes.Code)
}
