package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

// InitConfig 读取配置文件
func InitConfig() {
	// runtime.Caller(0) 用于获取当前源文件的路径，可以根据参数值获取不同调用栈层次的源文件路径。
	_, filePath, _, _ := runtime.Caller(0)

	currentDir := path.Dir(filePath)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(currentDir)

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
