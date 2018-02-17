package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sergey-shadrin/golang/gosite/model/database"
	"github.com/sergey-shadrin/golang/gosite/model/video"
)

type videoStatusInfo struct {
	Status status.VideoStatusValue `json:"status"`
}

func handleVideoStatus(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	contentKey, found := vars["CONTENT_KEY"]
	if found != true || contentKey == "" {
		log.Error("Invalid request method")
		http.Error(responseWriter, "Invalid request method", http.StatusBadRequest)
		return
	}

	q := "SELECT status FROM video WHERE content_key = ?"
	row := database.Get().QueryRow(q, contentKey)

	var videoStatusInfo videoStatusInfo;
	err := row.Scan(&videoStatusInfo.Status)
	if err != nil {
		http.Error(responseWriter, "The video not found", http.StatusNotFound)
		return
	}

	renderAsJson(responseWriter, videoStatusInfo)
}
