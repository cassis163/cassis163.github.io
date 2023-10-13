package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	components "github.com/cassis163/personal-site/components"
	"github.com/cassis163/personal-site/services"
)

type HomeHandler struct {
	BlogService *services.BlogService
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{
		BlogService: services.NewBlogService(),
	}
}

func (h *HomeHandler) Handle(c *gin.Engine) {
	c.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", components.HomePage())
	})
}
