package util

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
)

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		fmt.Printf("HTML: '%s'\n", html)
		return
	})
}
