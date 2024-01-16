package main

import (
	"video_service/config"
	"video_service/discovery"
	"video_service/logger"
	"video_service/pkg/db"
)

func main() {
	logger.InitLogger()      // 初始化日志
	config.InitConfig()      // 初始化配置文件
	db.InitDb()              // 初始化数据库
	db.InitRedis()           // 初始化缓存
	discovery.AutoRegister() // 自动注册
}
