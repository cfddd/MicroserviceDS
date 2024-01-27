package discovery

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"social_service/handler"
	social_server "social_service/server"
	"social_service/utils/etcd"
)

// AutoRegister etcd自动注册
func AutoRegister() {
	etcdAddress := viper.GetString("etcd.address")
	etcdRegister, err := etcd.NewEtcdRegister(etcdAddress)

	if err != nil {
		log.Fatal(err)
	}

	serviceName := viper.GetString("server.name")
	serviceAddress := viper.GetString("server.address")
	err = etcdRegister.ServiceRegister(serviceName, serviceAddress, 30)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	social_server.RegisterSocialServiceServer(server, handler.NewSocialService())

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
