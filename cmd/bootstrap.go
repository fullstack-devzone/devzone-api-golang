package cmd

import (
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/auth"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/posts"
	"github.com/sivaprasadreddy/devzone-api-golang/internal/api/users"
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

	userRep := domain.NewUserRepository(app.db)

	app.authController = auth.NewAuthController(config, userRep)

	app.userController = users.NewUserController(userRep)

	postsRepo := domain.NewPostRepo(app.db)
	app.postController = posts.NewPostController(postsRepo)

	return app
}
