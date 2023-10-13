package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	components "github.com/cassis163/personal-site/components"
	"github.com/cassis163/personal-site/services"
)

type BlogHandler struct {
	BlogService *services.BlogService
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{
		BlogService: services.NewBlogService(),
	}
}

func (h *BlogHandler) Handle(c *gin.Engine) {
	for _, blogPost := range h.BlogService.BlogPosts {
		c.GET("/blog/"+blogPost.FileName, func(c *gin.Context) {
			c.HTML(http.StatusOK, "hello.html", components.BlogPostPage())
		})
	}
}
