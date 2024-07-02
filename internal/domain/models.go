package domain

import (
	"time"
)

type UserModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

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
	CreatedBy   PostCreatedByModel `json:"createdBy"`
	CreatedDate *time.Time         `json:"createdAt"`
	UpdatedDate *time.Time         `json:"updatedAt"`
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
