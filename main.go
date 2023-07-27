package main

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/router"
	"log"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	router.SetupRouterGenre(r)

	err := r.Run()
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
