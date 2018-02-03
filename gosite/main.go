package main

import (
	"net/http"
	"github.com/sergey-shadrin/golang/gosite/handlers"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("/home/sergey/log/gosite.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Fatal(err)
	}
	defer file.Close()

	serverUrl := ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	router := handlers.Router()
	log.Fatal(http.ListenAndServe(serverUrl, router))
}