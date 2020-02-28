package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"log"
	proto "mini-micro-example/api/proto/info"
	"mini-micro-example/pkg/info"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.info"),
	)
	service.Init(micro.Action(func(context *cli.Context) {
		proto.RegisterEnvironmentHandler(service.Server(), &info.Info{})
	}), micro.AfterStart(func() error {
		return nil
	}), micro.AfterStop(func() error {
		return nil
	}))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
