package handlers

import (
	"net/http"
	"fmt"
)

func handleVideo(writer http.ResponseWriter, _ *http.Request) {
	response := `
{
  "id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
  "name": "Black Retrospetive Woman",
  "duration": 15,
  "thumbnail": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
  "url": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4"
}
`
	fmt.Fprint(writer, response)
}
