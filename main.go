package main

import (
	"github.com/gin-gonic/gin"
	"go_crud/initializers"
	"go_crud/router"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	router.SetupRouterGenre(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
