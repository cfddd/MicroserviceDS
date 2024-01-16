package oss7

import (
	"testing"
	"video_service/config"
	"video_service/logger"
)

// 测试上传初始化
func TestInitBucket(t *testing.T) {
	logger.InitLogger()
	config.InitConfig()
	InitBucket()
	//t.Log(QiniuClient.accessKey + "fgsdfgdfg")

	// 输出当前路径
	//t.Log(os.Getwd())

	err := UploadFile("./cfd", "cfd")
	if err != nil {
		logger.Log.Error(err)
	}
	return
}

// 测试获取视频地址
func TestGetFileUrl(t *testing.T) {
	logger.InitLogger()
	config.InitConfig()
	InitBucket()

	t.Log(GetFileUrl("douyin/video/v-16.mp4"))
}
