package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)

func logMiddleware(httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		httpHandler.ServeHTTP(w, r)
	})
}
