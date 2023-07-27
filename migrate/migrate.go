package main

import (
	"go_crud/initializers"
	"go_crud/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Film{},
		&models.Genre{},
		&models.User{},
	)
	if err != nil {
		panic(err)
	}
}
