package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/ur0/pokeintel-report/db"
	"github.com/ur0/pokeintel-report/router"
	"net/http"
	"time"
)

func main() {
	log.WithFields(log.Fields{"package": "main"}).Info("PokeIntel reports server starting")
	srv := &http.Server{
		Handler:      router.GetRouter(),
		Addr:         "0.0.0.0:8000", //TODO: Read this from somewhere
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	db.Connect()
	log.Fatal(srv.ListenAndServe())
}
