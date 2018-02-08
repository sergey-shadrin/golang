package handlers

import (
	"net/http"
)

type VideoInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Duration int `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	URL string `json:"url"`
}

func handleVideo(writer http.ResponseWriter, _ *http.Request) {
	video := VideoInfo{
		Id: "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name: "Black Retrospetive Woman",
		Duration: 15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		URL: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}

	renderAsJson(writer, video)
}
