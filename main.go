package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/config"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	//logFile := initLogging()
	//defer logFile.Close()
	initLogging()

	cfg := initConfig()
	app := initApp(cfg)
	router := setupRoutes(app)

	port := fmt.Sprintf(":%d", cfg.AppPort)
	srv := &http.Server{
		Handler:        router,
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("listening on port %d", cfg.AppPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func initApp(cfg config.Config) *App {
	app := NewApp(cfg)
	return app
}

func setupRoutes(app *App) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/links", app.LinkController.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/links/{id:[0-9]+}", app.LinkController.GetById).Methods(http.MethodGet)
	router.HandleFunc("/api/links", app.LinkController.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/links/{id:[0-9]+}", app.LinkController.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/links/{id:[0-9]+}", app.LinkController.Delete).Methods(http.MethodDelete)
	return router
}

func initLogging() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
}

func initFileLogging() *os.File {
	logFile, err := os.OpenFile("devzone.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	log.SetOutput(logFile)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	return logFile
}

func initConfig() config.Config {
	if _, err := os.Stat(".env"); errors.Is(err, fs.ErrNotExist) {
		log.Infof(".env file doesn't exist")
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Warningf("Couldn't load environment variables from .env file")
		}
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	AppPort, _ := strconv.Atoi(port)
	DbHost := os.Getenv("APP_DB_HOST")
	DbPort, _ := strconv.Atoi(os.Getenv("APP_DB_PORT"))
	DbUserName := os.Getenv("APP_DB_USERNAME")
	DbPassword := os.Getenv("APP_DB_PASSWORD")
	DbDatabase := os.Getenv("APP_DB_NAME")
	DbRunMigrations, _ := strconv.ParseBool(os.Getenv("APP_DB_RUN_MIGRATIONS"))
	return config.Config{
		AppPort:         AppPort,
		DbHost:          DbHost,
		DbPort:          DbPort,
		DbUserName:      DbUserName,
		DbPassword:      DbPassword,
		DbDatabase:      DbDatabase,
		DbRunMigrations: DbRunMigrations,
	}
}
