package discovery

import (
	"github.com/spf13/viper"
	"router_service/configs"
	"router_service/logger"
	"testing"
	"utils/etcd"
)

func TestDiscovery(t *testing.T) {
	logger.InitLogger()
	configs.InitConfig()
	etcdAddress := viper.GetString("etcd.address")
	serviceDiscovery, err := etcd.NewServiceDiscovery([]string{etcdAddress})

	if err != nil {
		logger.Log.Fatal(err)
	}
	defer serviceDiscovery.Close()
}
