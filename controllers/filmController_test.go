package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_crud/initializers"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func TestFilmCreate(t *testing.T) {

	initializers.DB = MakeTestDB(t)

	router := setupRouter()
	body := `{"FilmName": "Тестовый фильм", "ProductionYear": 2022, "Genres": "комедия,драма"}`
	req, _ := http.NewRequest("POST", "/films", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestFilmGetAll(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/films", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверка результата
	expectedResult := `{"films":null}`
	assert.Equal(t, expectedResult, strings.TrimSpace(w.Body.String()))
}
