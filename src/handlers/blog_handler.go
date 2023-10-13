package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
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
		html := h.BlogService.GenerateMdFromBlogPost(blogPost)
		component := Unsafe(string(html))
		c.GET("/blog/"+blogPost.FileName, func(c *gin.Context) {
			c.HTML(http.StatusOK, "blog_post.html", components.BlogPostPage(blogPost.Title, component))
		})
	}
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}
