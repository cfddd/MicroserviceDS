package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func InitRedis() {
	addr := viper.GetString("redis.address")
	pwd := viper.GetString("redis.password")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})
}
