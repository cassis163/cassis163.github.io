package main

import (
	"os"

	"github.com/gin-gonic/gin"

	handlers "github.com/cassis163/personal-site/handlers"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", gin.Dir("assets/images", false))
	r.HTMLRender = &TemplRender{}
	HomeHandler := handlers.NewHomeHandler()
	blogHandler := handlers.NewBlogHandler()
	HomeHandler.Handle(r)
	blogHandler.Handle(r)
	r.Run(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}
