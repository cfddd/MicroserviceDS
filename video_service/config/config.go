package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

// Config 数据结构
type Config struct {
	Mysql struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
	}
}

// ConfigData 配置数据变量
var ConfigData *Config

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

// DbDnsInit 拼接链接数据库的DNS
func DbDnsInit() string {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	InitConfig()
	dns := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8&parseTime=True&loc=Local"}, "")

	return dns
}

// InitRabbitMQUrl 拼接rabbitMQ连接地址
func InitRabbitMQUrl() string {
	user := viper.GetString("rabbitMQ.user")
	password := viper.GetString("rabbitMQ.password")
	address := viper.GetString("rabbitMQ.address")
	vhost := viper.GetString("rabbitMQ.vhost")
	url := strings.Join([]string{"amqp://", user, ":", password, "@", address, "/", vhost}, "")
	return url
}
