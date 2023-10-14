package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	components "github.com/cassis163/personal-site/components"
	"github.com/cassis163/personal-site/dtos"
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
	blogPostItems := getBlogPostItems(h)

	c.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", components.HomePage(blogPostItems))
	})
}

func getBlogPostItems(h *HomeHandler) []dtos.BlogPostPreviewItem {
	blogPostItems := []dtos.BlogPostPreviewItem{}
	blogPosts := h.BlogService.BlogPosts
	for _, blogPost := range blogPosts {
		formattedCreationDate := blogPost.CreatedAt.Format("January 2, 2006")
		link := fmt.Sprintf("/blog/%s", blogPost.FileName)
		readTime := h.BlogService.GetReadTimeInMinutes(blogPost)
		blogPostItems = append(blogPostItems, dtos.BlogPostPreviewItem{
			Title:        blogPost.Title,
			CreationDate: formattedCreationDate,
			ReadTime:     readTime,
			Link:         templ.SafeURL(link),
		})
	}
	return blogPostItems
}
