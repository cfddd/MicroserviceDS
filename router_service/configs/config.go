package configs

import (
	"github.com/spf13/viper"
	"path"
	"router_service/logger"
	"runtime"
)

// InitConfig 读取配置文件
func InitConfig() {
	_, filePath, _, _ := runtime.Caller(0)

	currentDir := path.Dir(filePath)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(currentDir)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	logger.Log.Info("读取配置文件--成功--")
}
