package discovery

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"router_service/logger"
	service "router_service/server"
	"router_service/utils/etcd"
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
	// 获取用户服务实例
	err = serviceDiscovery.ServiceDiscovery("user_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	userServiceAddr, _ := serviceDiscovery.GetService("user_service")
	userConn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("获取视频服务实例出错")
		logger.Log.Fatal(err)
	}
	userClient := service.NewUserServiceClient(userConn)
	logger.Log.Info("获取用户服务实例--成功--")
	instance["user_service"] = userClient

	// 获取视频服务实例
	err = serviceDiscovery.ServiceDiscovery("video_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	videoServiceAddr, _ := serviceDiscovery.GetService("video_service")
	videoConn, err := grpc.Dial(videoServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal(err)
	}

	videoClient := service.NewVideoServiceClient(videoConn)
	logger.Log.Info("获取视频服务实例--成功--")
	instance["video_service"] = videoClient

	// 获取社交服务实例
	err = serviceDiscovery.ServiceDiscovery("social_service")
	if err != nil {
		logger.Log.Fatal(err)
	}
	socialServiceAddr, _ := serviceDiscovery.GetService("social_service")
	socialConn, err := grpc.Dial(socialServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal(err)
	}

	socialClient := service.NewSocialServiceClient(socialConn)
	logger.Log.Info("获取社交服务实例--成功--")
	instance["social_service"] = socialClient

	return instance
}
