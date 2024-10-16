package main

import (
	"net/http"
	"router_service/configs"
	"router_service/discovery"
	"router_service/logger"
	"router_service/router"
	"time"
)

func main() {
	logger.InitLogger()
	configs.InitConfig()
	resolver := discovery.Resolver()
	r := router.InitRouter(resolver)
	server := &http.Server{
		Addr:           configs.GetServerAddr(),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		logger.Log.Warnln("启动失败...")
	}

}
