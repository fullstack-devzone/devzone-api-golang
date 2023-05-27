package posts

import (
	"time"

	"github.com/sivaprasadreddy/devzone-api-golang/users"
)

type Post struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Content     string     `json:"content"`
	CreatedBy   users.User `json:"created_by"`
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate time.Time  `json:"updated_date"`
}
