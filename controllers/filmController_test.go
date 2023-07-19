package controllers

import (
	"bytes"
	"go_crud/initializers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type FilmResponse struct {
	FilmName       string   `json:"FilmName"`
	ProductionYear int16    `json:"ProductionYear"`
	Genres         []string `json:"Genres"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Установка маршрутов
	r.POST("/films", FilmCreate)
	r.GET("/films", FilmGetAll)
	r.GET("/films/:id", GetFilmById)
	r.PUT("/films/:id", FilmUpDate)
	r.DELETE("/films/:id", FilmDelete)

	return r
}
func MakeTestDB(t *testing.T) *gorm.DB {
	db, err := initializers.InitTestDB("host=localhost user=postgres password=destro dbname=test_db_genre port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)

	}
	return db
}

var firstFilm string

func TestFilmCreate(t *testing.T) {

	initializers.DB = MakeTestDB(t)

	router := setupRouter()
	body := `{"FilmName": "Тестовый фильм", "ProductionYear": 2022, "Genres": "комедия,драма"}`
	req, _ := http.NewRequest("POST", "/films", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	firstFilm = w.Body.String()
}

func TestFilmGetAll(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/films", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, firstFilm, strings.TrimSpace(w.Body.String()))
}

func TestGetFilmById(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/film/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, firstFilm, strings.TrimSpace(w.Body.String()))

}

func TestFilmDelete(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/film/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResult := `{"film":"удален успешно"}`
	assert.Equal(t, expectedResult, strings.TrimSpace(w.Body.String()))
}

func TestFilmUpDate(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("PUT", "/film/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
