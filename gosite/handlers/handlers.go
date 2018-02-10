package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/list", handleList)
	subRouter.HandleFunc("/video/{VIDEO_ID}", handleVideo)
	return logMiddleware(router)
}
