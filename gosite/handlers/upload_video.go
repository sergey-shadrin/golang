package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"fmt"
	"crypto/md5"
)

func handleUploadVideo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		log.Error("Invalid request method")
		http.Error(responseWriter, "Invalid request method", http.StatusBadRequest)
		return
	}
	uploadedFile, fileHeader, err := request.FormFile("file[]")
	defer uploadedFile.Close()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType != "video/mp4" {
		http.Error(responseWriter, "Invalid content type", http.StatusBadRequest)
		return
	}

	h := md5.New()
	if _, err := io.Copy(h, uploadedFile); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	uploadedFile.Seek(0, 0)

	uploadedFileChecksum := string(h.Sum(nil))
	destinationDirName := fmt.Sprintf("/usr/local/www/data/gosite/content/%x", uploadedFileChecksum)
	if err := os.MkdirAll(destinationDirName, os.ModePerm); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	destinationFileName := fmt.Sprintf("%v/video.mp4", destinationDirName)
	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, uploadedFile); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
