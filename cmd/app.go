package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/auth"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/posts"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/users"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
)

type App struct {
	Router         *gin.Engine
	db             *pgx.Conn
	config         config.AppConfig
	userController *users.UserController
	postController *posts.PostController
	authController *auth.AuthenticationController
}

func NewApp(config config.AppConfig) *App {
	app := InitializeApp(config)
	app.Router = SetupRoutes(app)
	return app
}
