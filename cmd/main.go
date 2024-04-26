package main

import (
	"github.com/gin-gonic/gin"
	"gocrudchallange/app/handlers"
	"gocrudchallange/app/middleware"
)

func main() {
	r := gin.Default()
	middleware.SetupMiddleware(r)

	handlers.RegisterPersonRoutes(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
