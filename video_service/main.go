package main

import (
	"video_service/config"
	"video_service/model"
	"video_service/pkg/cache"
)

func main() {
	config.InitConfig() // 读取配文件
	model.InitDb()      // 初始化数据库
	cache.InitRedis()   // 初始化缓存

}
