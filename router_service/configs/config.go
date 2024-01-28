package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"router_service/logger"
	"runtime"
)

// InitConfig 读取配置文件
func InitConfig() {
	_, filePath, _, _ := runtime.Caller(0)
	currentDir := path.Dir(filePath)
	configPath := path.Join(currentDir, "config.yaml")

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	logger.Log.Info("读取配置文件--成功--")
}

func GetServerAddr() string {
	addr := viper.GetString("server.address")
	port := viper.GetString("server.port")

	return fmt.Sprintf("%s:%s", addr, port)
}
