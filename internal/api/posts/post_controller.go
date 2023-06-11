package posts

import (
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type PostController struct {
	repository domain.PostRepository
}

func NewPostController(repository domain.PostRepository) *PostController {
	return &PostController{repository}
}
