package posts

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserModel struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Roles       []string  `json:"roles"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type PostModel struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Content     string    `json:"content"`
	CreatedBy   UserModel `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type CreatePostModel struct {
	Title   string `json:"title" validate:"required"`
	Url     string `json:"url" validate:"required,url"`
	Content string `json:"content" validate:"required"`
}

func (l CreatePostModel) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Title, validation.Required),
		validation.Field(&l.Url, validation.Required, is.URL),
		validation.Field(&l.Content, validation.Required),
	)
}

type UpdatePostModel struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Content string `json:"content"`
}
