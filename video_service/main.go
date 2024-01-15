package main

import (
	"video_service/configs"
	"video_service/discovery"
	"video_service/init"
	"video_service/logger"
)

func main() {
	logger.InitLogger()      // 初始化日志
	configs.InitConfig()     // 初始化配置文件
	init.InitDb()            // 初始化数据库
	init.InitRedis()         // 初始化缓存
	discovery.AutoRegister() // 自动注册
}
