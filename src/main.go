package main

import (
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
	// Run on port 8080
	r.Run()
}
