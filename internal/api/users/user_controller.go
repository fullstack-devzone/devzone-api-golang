package users

import (
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type UserController struct {
	repository domain.UserRepository
}

func NewUserController(repository domain.UserRepository) *UserController {
	return &UserController{repository}
}
