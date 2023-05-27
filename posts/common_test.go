package posts_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sivaprasadreddy/devzone-api-golang/cmd"
	"github.com/sivaprasadreddy/devzone-api-golang/config"
	"github.com/sivaprasadreddy/devzone-api-golang/test_support"
)

var cfg config.AppConfig
var app *cmd.App
var router *gin.Engine

func TestMain(m *testing.M) {
	//Common Setup
	pgContainer := test_support.InitPostgresContainer()
	defer pgContainer.CloseFn()

	cfg = config.GetConfig(".env")
	app = cmd.NewApp(cfg)
	router = app.Router

	code := m.Run()

	//Common Teardown
	os.Exit(code)
}
