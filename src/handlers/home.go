package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	components "github.com/cassis163/personal-site/components"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Handle(c *gin.Engine) {
	c.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", components.HomePage())
	})
}
