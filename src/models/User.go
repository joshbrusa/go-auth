package models

import "time"

type User struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	ID        int
	Username  string
	Password  string
}
