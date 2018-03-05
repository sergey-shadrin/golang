package handlers

import (
	"net/http"
	"github.com/sergey-shadrin/golang/gosite/model/database"
)

type VideoListItem struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

func handleList(writer http.ResponseWriter, _ *http.Request) {
	q := "SELECT content_key, name, duration FROM video"
	rows, err := database.Get().Query(q)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var videoListItems []VideoListItem
	for rows.Next() {
		var videoListItem VideoListItem
		rows.Scan(&videoListItem.Id, &videoListItem.Name, &videoListItem.Duration)
		videoListItems = append(videoListItems, videoListItem)
	}
	renderAsJson(writer, videoListItems)
}
