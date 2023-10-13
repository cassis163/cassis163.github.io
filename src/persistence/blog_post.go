package persistence

import "time"

type BlogPost struct {
	FileName  string
	Title     string
	Content   []byte
	CreatedAt time.Time
}
