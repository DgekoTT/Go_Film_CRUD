package controllers_test

import (
	"go_crud/controllers"
	"go_crud/initializers"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouterGenre() *gin.Engine {
	r := gin.Default()
	r.POST("/genre", controllers.GenreCreate)
	r.GET("/genre", controllers.GenreGetAll)
	r.DELETE("/genre/:id", controllers.GenreDelete)
	return r
}
func MakeTestDB(t *testing.T) *gorm.DB {
	db, err := initializers.InitTestDB("host=localhost user=postgres password=destro dbname=test_db_genre port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)

	}
	return db
}
      bn  nn nnm mm
func