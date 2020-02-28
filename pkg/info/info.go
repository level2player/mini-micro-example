package info

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"log"
	proto "mini-micro-example/api/proto/info"
	"net"
	"os"
	"runtime"
)

type Info struct {
}

func (i *Info) GetEnvironmentInfo(ctx context.Context, req *proto.EnvironmentRequest, rsp *proto.EnvironmentResponse) error {
	user_id, ok := metadata.Get(ctx, "User_id")
	if !ok {
		return errors.BadRequest("mini.micro.example.info", "not found user_id")
	}
	log.Println(fmt.Sprintf("%v , visit environment", user_id))
	rsp.GoRoot = runtime.GOROOT()
	rsp.Version = runtime.Version()
	rsp.HostName, _ = os.Hostname()
	if ii, err := net.Interfaces(); err != nil {
		for idx, i := range ii {
			if idx < len(ii)-1 {
				rsp.NetInterface += i.Name + ","
			} else {
				rsp.NetInterface += i.Name
			}
		}
	}
	rsp.Goarch = runtime.GOARCH
	rsp.Goos = runtime.GOOS
	return nil
}
