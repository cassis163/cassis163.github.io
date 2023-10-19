package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	components "github.com/cassis163/personal-site/components"
	services "github.com/cassis163/personal-site/services"
	util "github.com/cassis163/personal-site/util"
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
	blogPostsHtmlByName, err := h.getBlogPostsHtmlByName()
	if err != nil {
		panic(err)
	}

	c.GET("/blog/:name", func(c *gin.Context) {
		name := c.Param("name")
		blogPost, err := h.BlogService.GetBlogPostByName(name)
		html := blogPostsHtmlByName[name]
		component := util.Unsafe(string(html))

		if err != nil {
			panic(err)
		}

		formattedCreationDate := blogPost.CreatedAt.Format("January 2, 2006")
		readTimeInMinutes := h.BlogService.GetReadTimeInMinutes(blogPost)
		c.HTML(http.StatusOK, "blog_post.html", components.BlogPostPage(blogPost.Title, formattedCreationDate, readTimeInMinutes, component))
	})
}

func (h *BlogHandler) getBlogPostsHtmlByName() (map[string]string, error) {
	var blogPostsHtmlByName = make(map[string]string)
	for _, blogPost := range h.BlogService.BlogPosts {
		html := h.BlogService.GenerateMdFromBlogPost(blogPost)
		blogPostsHtmlByName[blogPost.FileName] = string(html)
	}

	return blogPostsHtmlByName, nil
}
