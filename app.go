package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/config"
	"github.com/sivaprasadreddy/devzone-api-golang/database"
	"github.com/sivaprasadreddy/devzone-api-golang/links"
	"net/http"
	"os"
)

type App struct {
	Router         *mux.Router
	db             *sql.DB
	linkController *links.LinkController
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

	linksRepo := links.NewLinkRepo(app.db)
	linkSvc := links.NewLinkService(linksRepo)
	app.linkController = links.NewLinkController(linkSvc)

	app.Router = app.setupRoutes()
}

func (app *App) setupRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()
	app.setupLinkRoutes(apiRouter, app.linkController)
	return router
}

func (app *App) setupLinkRoutes(router *mux.Router, controller *links.LinkController) {
	r := router.PathPrefix("/links").Subrouter()
	r.HandleFunc("", controller.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/{id:[0-9]+}", controller.GetById).Methods(http.MethodGet)
	r.HandleFunc("", controller.Create).Methods(http.MethodPost)
	r.HandleFunc("/{id:[0-9]+}", controller.Update).Methods(http.MethodPut)
	r.HandleFunc("/{id:[0-9]+}", controller.Delete).Methods(http.MethodDelete)
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
