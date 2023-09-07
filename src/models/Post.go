package models

import "time"

type Post struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	Id        int
	Title     string
	Body      string
}
