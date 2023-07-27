package main

import (
	"go_crud/initializers"
	"go_crud/router"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	router.SetupRouterGenre(r)
	router.SetupAuthRouter(r)

	err := r.Run()
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
