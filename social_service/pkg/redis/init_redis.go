package redis

import (
	"fmt"
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

	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		//global.Log.Errorf("Redis 连接失败 %s", err) //TODO: 连接失败日志
		panic(err)
	}

	fmt.Println("Redis 连接成功!")
	//global.Log.Info("Redis 连接成功!") //TODO: 连接成功日志
}
