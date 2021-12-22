package links

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
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

type LinkModel struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Tags        []string  `json:"tags"`
	CreatedBy   UserModel `json:"created_by"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type CreateLinkModel struct {
	Title string   `json:"title" validate:"required"`
	Url   string   `json:"url" validate:"required,url"`
	Tags  []string `json:"tags"`
}

func (l CreateLinkModel) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Title, validation.Required),
		validation.Field(&l.Url, validation.Required, is.URL),
	)
}

type UpdateLinkModel struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Url   string   `json:"url"`
	Tags  []string `json:"tags"`
}
