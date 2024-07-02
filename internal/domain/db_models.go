package domain

import (
	"time"
)

type Post struct {
	Id          int
	Title       string
	Url         string
	Content     string
	CreatedBy   int
	CreatedDate *time.Time
	UpdatedDate *time.Time
}

type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	Role        string
	CreatedDate *time.Time
	UpdatedDate *time.Time
}
