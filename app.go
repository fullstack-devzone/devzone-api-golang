package main

import (
	"github.com/sivaprasadreddy/devzone-api-golang/config"
	"github.com/sivaprasadreddy/devzone-api-golang/database"
	"github.com/sivaprasadreddy/devzone-api-golang/links"
)

type App struct {
	LinkController *links.LinkController
}

func NewApp(config config.Config) *App {
	app := &App{}
	app.init(config)
	return app
}

func (app *App) init(config config.Config) {
	db := database.GetDb(config)

	linksRepo := links.NewLinkRepo(db)
	linkSvc := links.NewLinkService(linksRepo)
	app.LinkController = links.NewLinkController(linkSvc)
}
