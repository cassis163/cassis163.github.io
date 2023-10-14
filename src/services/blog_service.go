package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	dtos "github.com/cassis163/personal-site/dtos"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const MARKDOWN_EXTENSION = ".md"

// In words per minute
const AVERAGE_READING_SPEED = 200

func NewBlogService() *BlogService {
	return &BlogService{
		BlogPosts: getBlogPosts(),
	}
}

type BlogService struct {
	BlogPosts []dtos.BlogPost
}

func (b *BlogService) GenerateMdFromBlogPost(blogPost dtos.BlogPost) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(blogPost.Content)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func (b *BlogService) GetBlogPostByName(name string) (dtos.BlogPost, error) {
	for _, blogPost := range b.BlogPosts {
		if blogPost.FileName == name {
			return blogPost, nil
		}
	}

	return dtos.BlogPost{}, errors.New("blog post not found")
}

func (b *BlogService) GetReadTimeInMinutes(blogPost dtos.BlogPost) string {
	words := len(strings.Fields(string(blogPost.Content)))
	// read time in minutes
	readTime := words / AVERAGE_READING_SPEED
	return fmt.Sprintf("%dm", readTime)
}

func getBlogPosts() []dtos.BlogPost {
	dir := getBlogPostsPath()
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var blogPosts []dtos.BlogPost
	for _, file := range files {
		if filepath.Ext(file.Name()) != MARKDOWN_EXTENSION {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			panic(err)
		}

		fileName := strings.TrimSuffix(file.Name(), MARKDOWN_EXTENSION)
		blogPosts = append(blogPosts, dtos.BlogPost{
			FileName:  fileName,
			Title:     convertFileNameToTitle(fileName),
			Content:   content,
			CreatedAt: fileInfo.ModTime(),
		})
	}

	return blogPosts
}

func convertFileNameToTitle(fileName string) string {
	fileName = strings.ReplaceAll(fileName, "-", " ")
	caser := cases.Title(language.English)
	return caser.String(fileName)
}

func getBlogPostsPath() string {
	wdPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	blogPostsDir := filepath.Join(wdPath, "assets", "blog-posts")
	if _, err := os.Stat(blogPostsDir); os.IsNotExist(err) {
		panic("Assets directory not found.")
	}

	return blogPostsDir
}
