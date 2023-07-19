package main

import (
	"go_crud/initializers"
	"go_crud/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
	initializers.InitTestDB("host=localhost user=postgres password=destro dbname=test_db_genre port=5433 sslmode=disable")
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Film{}, &models.Genre{})
	err2 := initializers.TestDB.AutoMigrate(&models.Film{}, &models.Genre{})
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
}
