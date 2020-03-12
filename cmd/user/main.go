package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"log"
	proto "mini-micro-example/api/proto/user"
	"mini-micro-example/pkg/user"
)

//main3
func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.user"),
	)
	service.Init(micro.Action(func(context *cli.Context) {
		proto.RegisterUserHandler(service.Server(), &user.User{})
	}), micro.AfterStart(func() error {
		return nil
	}), micro.AfterStop(func() error {
		return nil
	}))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
