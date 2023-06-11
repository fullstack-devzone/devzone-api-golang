package domain

import (
	"time"
)

type PostCreatedByModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostModel struct {
	Id          int                `json:"id"`
	Title       string             `json:"title"`
	Url         string             `json:"url"`
	Content     string             `json:"content"`
	CreatedBy   PostCreatedByModel `json:"created_by"`
	CreatedDate *time.Time         `json:"created_date"`
	UpdatedDate *time.Time         `json:"updated_date"`
}
