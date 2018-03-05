package handlers

import (
	"net/http"
	"github.com/sergey-shadrin/golang/gosite/model/database"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sergey-shadrin/golang/gosite/path_generator"
)

type VideoInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
}

func handleVideo(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	contentKey, found := vars["CONTENT_KEY"]
	if found != true || contentKey == "" {
		log.Error("Invalid request method")
		http.Error(responseWriter, "Invalid request method", http.StatusBadRequest)
		return
	}

	q := "SELECT content_key, name, duration FROM video WHERE content_key = ?"
	row := database.Get().QueryRow(q, contentKey)

	var videoInfo VideoInfo
	err := row.Scan(&videoInfo.Id, &videoInfo.Name, &videoInfo.Duration)
	if err != nil {
		log.Error(err.Error())
		http.Error(responseWriter, "The video not found", http.StatusNotFound)
		return
	}

	videoInfo.URL = path_generator.GetVideoUrl(videoInfo.Id)
	videoInfo.Thumbnail = path_generator.GetVideoThumbnailUrl(videoInfo.Id)

	renderAsJson(responseWriter, videoInfo)
}
