package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
)

type App struct {
	Router         *gin.Engine
	db             *pgx.Conn
	config         config.AppConfig
	postController *api.PostController
	authController *api.AuthHandler
}

func NewApp(config config.AppConfig) *App {
	app := InitializeApp(config)
	app.Router = SetupRoutes(app)
	return app
}
