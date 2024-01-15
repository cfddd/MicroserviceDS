package main

import (
	"user_service/config"
	"user_service/internal/model"
	"user_service/pkg/redis"
)

func main() {
	config.InitConfig() //初始化配置文件
	redis.InitRedis()   //初始化redis
	model.InitDb()      //初始化数据库
}
