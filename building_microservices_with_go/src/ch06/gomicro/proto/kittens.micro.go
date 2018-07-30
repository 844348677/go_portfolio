// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: kittens.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	kittens.proto

It has these top-level messages:
	RequestEnvelope
	ResponseEnvelope
	Request
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Kittens service

type KittensService interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type kittensService struct {
	c    client.Client
	name string
}

func NewKittensService(name string, c client.Client) KittensService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &kittensService{
		c:    c,
		name: name,
	}
}

func (c *kittensService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Kittens.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Kittens service

type KittensHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterKittensHandler(s server.Server, hdlr KittensHandler, opts ...server.HandlerOption) error {
	type kittens interface {
		Hello(ctx context.Context, in *Request, out *Response) error
	}
	type Kittens struct {
		kittens
	}
	h := &kittensHandler{hdlr}
	return s.Handle(s.NewHandler(&Kittens{h}, opts...))
}

type kittensHandler struct {
	KittensHandler
}

func (h *kittensHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.KittensHandler.Hello(ctx, in, out)
}
