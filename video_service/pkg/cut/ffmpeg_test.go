package cut

import (
	"github.com/google/uuid"
	"testing"
	"video_service/config"
	"video_service/logger"
	"video_service/pkg/oss7"
)

func TestCover(t *testing.T) {
	logger.InitLogger()
	config.InitConfig()
	oss7.InitBucket()
	videoURL := "http://douyin.cfddfc.online/douyin/video/cfd--3be9265a-309e-499b-b371-ce867e2cc60a.mp4"
	imageBytes, _ := Cover(videoURL, "00:00:01")
	UUID := uuid.New()
	pictureDir := "douyin/cover/" + UUID.String() + ".jpg"
	oss7.UploadFileWithByte(pictureDir, imageBytes)
}
