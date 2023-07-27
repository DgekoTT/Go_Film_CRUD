package controllers_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"go_crud/initializers"
	"go_crud/models"
	"go_crud/router"
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

var (
	FilmRouter  *gin.Engine
	GenreRouter *gin.Engine
	r           *gin.Engine
	firstFilm   string
	t           *testing.T
)

func init() {
	recreateTestDatabase()
	r = gin.Default()
	initializers.DB = MakeTestDB(t)
	MigrationTestDB()
	FilmRouter = router.SetupRouter(r)
	GenreRouter = router.SetupRouterGenre(r)
	// Тут также можно добавить другие инициализации, если они необходимы
}

type FilmResponse struct {
	FilmName       string   `json:"FilmName"`
	ProductionYear int16    `json:"ProductionYear"`
	Genres         []string `json:"Genres"`
}

func MakeTestDB(t *testing.T) *gorm.DB {
	db, err := initializers.InitTestDB("host=localhost user=postgres password=destro dbname=test_db_genre port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)

	}
	err2 := initializers.TestDB.AutoMigrate(&models.Film{}, &models.Genre{})
	if err2 != nil {
		log.Println("Failed to migrate")
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
	dbName := "test_db_genre"
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

func TestFilmCreate(t *testing.T) {

	genres := []string{"комедия", "драма"}
	for _, genre := range genres {
		genreBody := fmt.Sprintf(`{"genre_name": "%s"}`, genre)
		genreReq, _ := http.NewRequest("POST", "/genre", bytes.NewBufferString(genreBody))
		genreReq.Header.Set("Content-Type", "application/json")
		genreRes := httptest.NewRecorder()
		GenreRouter.ServeHTTP(genreRes, genreReq)

		assert.Equal(t, http.StatusOK, genreRes.Code)
	}
	body := `{"FilmName": "Тестовый фильм", "ProductionYear": 2022, "Genres": "комедия,драма"}`
	req, _ := http.NewRequest("POST", "/film", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	FilmRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	firstFilm = w.Body.String()
}

func TestFilmGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/films", nil)
	w := httptest.NewRecorder()
	FilmRouter.ServeHTTP(w, req)
	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, w.Code)

	//assert.Equal(t, firstFilm, strings.TrimSpace(w.Body.String()))
}

func TestGetFilmById(t *testing.T) {
	req, _ := http.NewRequest("GET", "/film/id/1", nil)
	w := httptest.NewRecorder()
	FilmRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, firstFilm, strings.TrimSpace(w.Body.String()))

}

func TestFilmUpDate(t *testing.T) {
	jsonData := `{"FilmName": "Новое имя фильма", "ProductionYear": 2023}`
	req, _ := http.NewRequest("PUT", "/film/id/1", strings.NewReader(jsonData))
	w := httptest.NewRecorder()
	FilmRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestFilmDelete(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/film/id/1", nil)
	w := httptest.NewRecorder()
	FilmRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedResult := `{"film":"удален успешно"}`
	assert.Equal(t, expectedResult, strings.TrimSpace(w.Body.String()))
}
