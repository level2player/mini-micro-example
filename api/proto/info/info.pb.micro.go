// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: info/info.proto

package info

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Environment service

type EnvironmentService interface {
	GetEnvironmentInfo(ctx context.Context, in *EnvironmentRequest, opts ...client.CallOption) (*EnvironmentResponse, error)
}

type environmentService struct {
	c    client.Client
	name string
}

func NewEnvironmentService(name string, c client.Client) EnvironmentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "environment"
	}
	return &environmentService{
		c:    c,
		name: name,
	}
}

func (c *environmentService) GetEnvironmentInfo(ctx context.Context, in *EnvironmentRequest, opts ...client.CallOption) (*EnvironmentResponse, error) {
	req := c.c.NewRequest(c.name, "Environment.GetEnvironmentInfo", in)
	out := new(EnvironmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Environment service

type EnvironmentHandler interface {
	GetEnvironmentInfo(context.Context, *EnvironmentRequest, *EnvironmentResponse) error
}

func RegisterEnvironmentHandler(s server.Server, hdlr EnvironmentHandler, opts ...server.HandlerOption) error {
	type environment interface {
		GetEnvironmentInfo(ctx context.Context, in *EnvironmentRequest, out *EnvironmentResponse) error
	}
	type Environment struct {
		environment
	}
	h := &environmentHandler{hdlr}
	return s.Handle(s.NewHandler(&Environment{h}, opts...))
}

type environmentHandler struct {
	EnvironmentHandler
}

func (h *environmentHandler) GetEnvironmentInfo(ctx context.Context, in *EnvironmentRequest, out *EnvironmentResponse) error {
	return h.EnvironmentHandler.GetEnvironmentInfo(ctx, in, out)
}
