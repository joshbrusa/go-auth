package models

import "time"

type Post struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	ID        int
	Title     string
	Body      string
}
