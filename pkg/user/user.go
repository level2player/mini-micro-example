package user

import (
	"context"
	"github.com/micro/go-micro/errors"
	proto "mini-micro-example/api/proto/user"
	"mini-micro-example/internal/pkg"
)

type User struct {
	jwter pkg.Jwter
}

func (u *User) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	if req.UserId != "jacky" || req.Password != "123456" {
		return errors.BadRequest("mini.micro.example.user", "user_id or password error")
	}
	if token, err := u.jwter.GetUserNewJwt(req.UserId); err != nil {
		return errors.BadRequest("mini.micro.example.user", "build jwt error")
	} else {
		rsp.Token = token
	}
	return nil
}
