package main

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/cassis163/personal-site/handlers"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", gin.Dir("../assets/images", false))
	r.HTMLRender = &TemplRender{}
	HomeHandler := handlers.NewHomeHandler()
	blogHandler := handlers.NewBlogHandler()
	HomeHandler.Handle(r)
	blogHandler.Handle(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
