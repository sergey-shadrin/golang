package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/list", handleList)
	subRouter.HandleFunc("/video", handleUploadVideo)
	subRouter.HandleFunc("/video/{CONTENT_KEY}", handleVideo)
	subRouter.HandleFunc("/video/{CONTENT_KEY}/status", handleVideoStatus)
	return logMiddleware(router)
}
