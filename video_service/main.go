package main

import (
	"video_service/config"
	"video_service/discovery"
	"video_service/handler"
	"video_service/logger"
	"video_service/pkg/cache"
	"video_service/pkg/db"
	"video_service/pkg/oss7"
)

func main() {
	logger.InitLogger() // 初始化日志
	config.InitConfig() // 初始化配置文件
	db.InitDb()         // 初始化数据库
	cache.InitRedis()   // 初始化缓存
	oss7.InitBucket()   // 初始化OSS

	go func() {
		handler.PublishVideo()
	}()

	discovery.AutoRegister() // 自动注册
}
