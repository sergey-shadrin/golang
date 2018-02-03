package main

import (
	"net/http"
	"github.com/sergey-shadrin/golang/gosite/handlers"
	log "github.com/sirupsen/logrus"
	"os"
)

/*
func main() {
	http.HandleFunc("/api/v1/list", func(writer http.ResponseWriter, _ *http.Request) {
		response := `
[
  {
    "id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
    "name": "Black Retrospetive Woman",
    "duration": 15,
    "thumbnail": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg"
  }
]`
		fmt.Fprint(writer, response)
	})

	http.HandleFunc("/api/v1/video/d290f1ee-6c54-4b01-90e6-d701748f0851", func(writer http.ResponseWriter, _ *http.Request) {
		response := `
{
  "id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
  "name": "Black Retrospetive Woman",
  "duration": 15,
  "thumbnail": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
  "url": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4"
}
`
		fmt.Fprint(writer, response)
	})
	http.ListenAndServe(":8000", nil)
}
*/

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