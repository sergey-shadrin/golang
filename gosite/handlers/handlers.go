package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/list", handleList)
	subRouter.HandleFunc("/video/d290f1ee-6c54-4b01-90e6-d701748f0851", handleVideo)
	return logMiddleware(router)
}