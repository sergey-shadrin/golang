package path_generator

import "fmt"

func GetVideoUrl(contentKey string) string {
	return fmt.Sprintf("/content/%s/video.mp4", contentKey)
}

func GetVideoThumbnailUrl(contentKey string) string {
	return fmt.Sprintf("/content/%s/screen.jpg", contentKey)
}