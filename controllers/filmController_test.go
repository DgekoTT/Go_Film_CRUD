package controllers_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"go_crud/controllers"
	"go_crud/initializers"
	"go_crud/models"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	r.POST("/films", controllers.FilmCreate)
	r.POST("/genre", controllers.GenreCreate)
	r.GET("/films", controllers.FilmGetAll)
	r.GET("/films/id/:id", controllers.GetFilmById)
	//r.PUT("/films/:id", controllers.FilmUpDate)
	r.DELETE("/films/id/:id", controllers.FilmDelete)

	return r
}
func MakeTestDB(t *testing.T) *gorm.DB {
	db, err := initializers.InitTestDB("host=localhost user=postgres password=destro dbname=test_db_genre port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)

	}
	err2 := initializers.TestDB.AutoMigrate(&models.Film{}, &models.Genre{})
	if err2 != nil {
		panic(err2)
	}
	return db
}

func MigrationTestDB() {
	err2 := initializers.TestDB.AutoMigrate(&models.Film{}, &models.Genre{})
	if err2 != nil {
		panic(err2)
	}
}

func recreateTestDatabase() {
	dbName := "test_db_genre" // Замените на имя вашей тестовой базы данных
	db, err := sql.Open("postgres", "postgres://postgres:destro@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка при подключении к PostgreSQL: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
	if err != nil {
		log.Fatalf("Ошибка при удалении базы данных: %v", err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		log.Fatalf("Ошибка при создании базы данных: %v", err)
	}

	fmt.Println("Тестовая база данных создана успешно.")
}

var firstFilm string

func TestFilmCreate(t *testing.T) {
	recreateTestDatabase()
	initializers.DB = MakeTestDB(t)
	MigrationTestDB()
	router := setupRouter()

	genres := []string{"комедия", "драма"}
	for _, genre := range genres {
		genreBody := fmt.Sprintf(`{"genre_name": "%s"}`, genre)
		genreReq, _ := http.NewRequest("POST", "/genre", bytes.NewBufferString(genreBody))
		genreReq.Header.Set("Content-Type", "application/json")
		genreRes := httptest.NewRecorder()
		router.ServeHTTP(genreRes, genreReq)

		assert.Equal(t, http.StatusOK, genreRes.Code)
	}
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

	//assert.Equal(t, firstFilm, strings.TrimSpace(w.Body.String()))
}

func TestGetFilmById(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/films/id/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, firstFilm, strings.TrimSpace(w.Body.String()))

}

//func TestFilmUpDate(t *testing.T) {
//	initializers.DB = MakeTestDB(t)
//	router := setupRouter()
//	jsonData := `{"FilmName": "Новое имя фильма", "ProductionYear": 2023}`
//	req, _ := http.NewRequest("PUT", "/films/1", strings.NewReader(jsonData))
//	w := httptest.NewRecorder()
//	router.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//
//}

func TestFilmDelete(t *testing.T) {
	initializers.DB = MakeTestDB(t)
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/films/id/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResult := `{"film":"удален успешно"}`
	assert.Equal(t, expectedResult, strings.TrimSpace(w.Body.String()))
}
