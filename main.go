package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sivaprasadreddy/devzone-api-golang/cmd"
	"github.com/sivaprasadreddy/devzone-api-golang/config"
)

func main() {
	cfg := config.GetConfig(".env")
	app := cmd.NewApp(cfg)

	port := fmt.Sprintf(":%d", cfg.AppPort)
	srv := &http.Server{
		Handler:        app.Router,
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
