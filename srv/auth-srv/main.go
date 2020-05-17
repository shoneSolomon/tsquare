package main

import (
	"auth-srv/handler"
	"auth-srv/model"
	"fmt"

	"github.com/micro/cli"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"auth-srv/basic"
	"auth-srv/basic/config"
	auth "auth-srv/proto/auth"
)

func main() {
	// 初始化配置
	basic.Init()

	// 注册服务
	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) error {
			// 初始化model
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	// Register Handler
	_ = auth.RegisterServiceHandler(service.Server(), new(handler.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
