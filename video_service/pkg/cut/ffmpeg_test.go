package cut

import (
	"testing"
	"video_service/config"
	"video_service/logger"
	"video_service/pkg/oss7"
)

func TestCover(t *testing.T) {
	logger.InitLogger()
	config.InitConfig()
	oss7.InitBucket()
	videoURL := "http://douyin.cfddfc.online/douyin/video/v-16.mp4?e=1705413504&token=1gvutSKAKY7cwPwwWkHDeQL8kX4iCuOxGvTMg2CT:vG2jKYpY_rTFmeIf3eb3YfTHRWk="

	imageBytes, _ := Cover(videoURL, "00:00:05")

	oss7.UploadFileWithByte("output.jpg", imageBytes)
}
