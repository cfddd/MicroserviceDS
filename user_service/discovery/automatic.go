package discovery

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_service/internal/handler"
	"user_service/server"
	"utils/etcd"
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

	newserver := grpc.NewServer()
	server.RegisterUserServiceServer(newserver, handler.NewUserService())

	err = newserver.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
