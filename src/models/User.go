package models

import "time"

type User struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	Id        int
	Username  string
	Password  string
}
