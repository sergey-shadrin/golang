package main

import (
	"net/http"
	"github.com/sergey-shadrin/golang/gosite/handlers"
	log "github.com/sirupsen/logrus"
	"os"
	"context"
	"github.com/sergey-shadrin/golang/gosite/osutil"
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

	killSignalChan := osutil.GetKillSignalChan()
	server := startServer(":8000")

	osutil.WaitForKillSignal(killSignalChan)
	server.Shutdown(context.Background())
}

func startServer(serverUrl string) *http.Server {
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")
	router := handlers.Router()
	server := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatal(http.ListenAndServe(serverUrl, router))
	}()
	return server
}
