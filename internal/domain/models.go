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

type PostsPageModel struct {
	TotalElements int         `json:"totalElements"`
	TotalPages    int         `json:"totalPages"`
	PageNumber    int         `json:"pageNumber"`
	IsFirst       bool        `json:"isFirst"`
	IsLast        bool        `json:"isLast"`
	HasNext       bool        `json:"hasNext"`
	HasPrevious   bool        `json:"hasPrevious"`
	Data          []PostModel `json:"data"`
}
