package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"fmt"
	"github.com/sergey-shadrin/golang/gosite/model/database"
	"github.com/sergey-shadrin/golang/gosite/key_generator"
	"github.com/sergey-shadrin/golang/gosite/model/video"
	"mime/multipart"
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

	db := database.Get()
	q := "INSERT INTO video set status = ?"
	r, err := db.Exec(q, status.INIT)
	if err != nil {
		handleInternalError(responseWriter, err)
		return
	}

	var lastInsertId int64;
	if lastInsertId, err = r.LastInsertId(); err != nil {
		handleInternalError(responseWriter, err)
		return
	}

	var contentKey string;
	if contentKey, err = key_generator.GenerateKeyById(lastInsertId); err != nil {
		handleInternalError(responseWriter, err)
		return
	}

	if err = copyUploadedFileToStorage(uploadedFile, contentKey); err != nil {
		handleInternalError(responseWriter, err)
		return
	}

	q = "UPDATE video SET content_key = ? WHERE id = ?"
	if _, err = db.Exec(q, contentKey, lastInsertId); err != nil {
		handleInternalError(responseWriter, err)
	}
}

func handleInternalError(responseWriter http.ResponseWriter, err error) {
	http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
}

func copyUploadedFileToStorage(uploadedFile multipart.File, contentKey string) error {
	destinationDirName := fmt.Sprintf("/usr/local/www/data/gosite/content/%x", contentKey)
	if err := os.MkdirAll(destinationDirName, os.ModePerm); err != nil {
		return err
	}

	destinationFileName := fmt.Sprintf("%v/video.mp4", destinationDirName)
	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, uploadedFile); err != nil {
		return err
	}

	return nil
}