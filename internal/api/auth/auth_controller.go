package auth

import (
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

type AuthenticationController struct {
	cfg        config.AppConfig
	repository domain.UserRepository
}

func NewAuthController(cfg config.AppConfig, repository domain.UserRepository) *AuthenticationController {
	return &AuthenticationController{
		cfg:        cfg,
		repository: repository,
	}
}
