package domain

import (
	"time"
)

type Post struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Content     string     `json:"content"`
	CreatedBy   int        `json:"created_by"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date"`
}

type User struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Role        string     `json:"role"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date"`
}
