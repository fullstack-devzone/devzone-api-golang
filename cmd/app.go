package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/config"
	"github.com/sivaprasadreddy/devzone-api-golang/database"
	"github.com/sivaprasadreddy/devzone-api-golang/posts"
)

type App struct {
	Router         *gin.Engine
	db             *pgx.Conn
	postController *posts.PostController
}

func NewApp(config config.AppConfig) *App {
	app := &App{}
	app.init(config)
	return app
}

func (app *App) init(config config.AppConfig) {
	//logFile := initLogging()
	//defer logFile.Close()
	app.initLogging()

	app.db = database.GetDb(config)

	postsRepo := posts.NewPostRepo(app.db)
	app.postController = posts.NewPostController(postsRepo)

	app.Router = app.setupRoutes()
}

func (app *App) setupRoutes() *gin.Engine {
	r := gin.Default()
	apiRouter := r.Group("/api/posts")
	{
		apiRouter.GET("", app.postController.GetAll)
		apiRouter.GET("/:id", app.postController.GetById)
		apiRouter.POST("", app.postController.Create)
		apiRouter.PUT("/:id", app.postController.Update)
		apiRouter.DELETE("/:id", app.postController.Delete)
	}
	return r
}

func (app *App) initLogging() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
}

func (app *App) initFileLogging() *os.File {
	logFile, err := os.OpenFile("devzone.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	return logFile
}
