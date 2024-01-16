package main

import (
	"social_service/config"
	"social_service/model"
	"social_service/pkg/redis"
)

func main() {
	config.InitConfig() // 读取配置文件
	redis.InitRedis()   //初始化 redis
	model.InitDb()      //初始化 mysql
}
