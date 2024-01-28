package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

type Config struct {
	Mysql struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
	}
}

// InitConfig 读取配置文件
func InitConfig() {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filePath)
	configPath := path.Join(currentDir, "config.yaml")

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// DbDsnInit 连接数据库的DSN
func DbDsnInit() string {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")

	InitConfig()
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8&parseTime=True&loc=Local"}, "")

	return dsn
}

// RedisURlInit 连接redis
func RedisURlInit() string {
	password := viper.GetString("redis.password")
	address := viper.GetString("redis.address")
	database := viper.GetString("redis.database")

	InitConfig()
	RedisURL := "redis://" + password + "@" + address + "/" + database

	return RedisURL
}
