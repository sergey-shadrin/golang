package handlers

import (
	"net/http"
	"io"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

func renderAsJson(writer http.ResponseWriter, val interface{}) {
	encodedVideoListItem, err := json.Marshal(val)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(writer, string(encodedVideoListItem)); err != nil {
		log.WithField("err", err).Error("write response error")
	}

}
