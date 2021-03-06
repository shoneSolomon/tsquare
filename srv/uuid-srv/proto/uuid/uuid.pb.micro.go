// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: uuid.proto

package go_micro_service_uuid

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Uuid service

func NewUuidEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Uuid service

type UuidService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type uuidService struct {
	c    client.Client
	name string
}

func NewUuidService(name string, c client.Client) UuidService {
	return &uuidService{
		c:    c,
		name: name,
	}
}

func (c *uuidService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Uuid.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Uuid service

type UuidHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterUuidHandler(s server.Server, hdlr UuidHandler, opts ...server.HandlerOption) error {
	type uuid interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type Uuid struct {
		uuid
	}
	h := &uuidHandler{hdlr}
	return s.Handle(s.NewHandler(&Uuid{h}, opts...))
}

type uuidHandler struct {
	UuidHandler
}

func (h *uuidHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UuidHandler.Call(ctx, in, out)
}
