package cmd

import (
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/config"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/db"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/domain"
)

func InitializeApp(config config.AppConfig) *App {
	app := &App{}
	app.config = config
	app.InitLogging()
	//logFile := app.InitFileLogging()
	//defer logFile.Close()

	app.db = db.GetDb(config)

	postsRepo := domain.NewPostRepo(app.db)
	app.postController = api.NewPostController(postsRepo)

	userRep := domain.NewUserRepository(app.db)
	app.authController = api.NewAuthHandler(userRep)

	return app
}
