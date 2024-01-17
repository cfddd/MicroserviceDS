package discovery

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// todo
	// 获取用户服务实例
	err = serviceDiscovery.ServiceDiscovery("user_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	userServiceAddr, _ := serviceDiscovery.GetService("user_service")
	_, err = grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("获取视频服务实例出错")
		logger.Log.Fatal(err)
	}
}
