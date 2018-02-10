package handlers

import (
	"net/http"
)

type VideoListItem struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

func handleList(writer http.ResponseWriter, _ *http.Request) {
	videoListItem := VideoListItem{
		Id:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospetive Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
	}
	renderAsJson(writer, [1]VideoListItem{videoListItem})
}
