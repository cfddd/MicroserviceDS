package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// 连接redis
var Redis *redis.Client

var Ctx = context.Background()

func InitRedis() {
	addr := viper.GetString("redis.address")
	pwd := viper.GetString("redis.password")
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0, // 存入DB0
	})
}
