package dtos

import "github.com/a-h/templ"

type BlogPostPreviewItem struct {
	Title        string
	CreationDate string
	ReadTime     string
	Link         templ.SafeURL
}
