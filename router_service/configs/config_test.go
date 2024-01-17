package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"router_service/logger"
	"testing"
)

func TestConfig(t *testing.T) {
	logger.InitLogger()
	InitConfig()
	fmt.Println(viper.GetString("server.name"))
	fmt.Println(viper.GetString("server.port"))
	fmt.Println(viper.GetString("server.address"))
	fmt.Println(viper.GetString("etcd.address"))

}
