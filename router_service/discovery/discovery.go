package discovery

import (
	"github.com/spf13/viper"
	"router_service/logger"
	"utils/etcd"
)

func Resolver() map[string]interface{} {
	instance := make(map[string]interface{})

	etcdAddress := viper.GetString("etcd.address")
	serviceDiscovery, err := etcd.NewServiceDiscovery([]string{etcdAddress})

	if err != nil {
		logger.Log.Fatal(err)
	}
	defer serviceDiscovery.Close()

	// todo

	return instance
}
