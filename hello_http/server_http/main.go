package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	gw "CoolCar/proto/hello_http"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// type HelloHTTPServer interface {
//定义SayHello方法
// SayHello(context.Context, *HelloHTTPRequest) (*HelloHTTPResponse, error)
// mustEmbedUnimplementedHelloHTTPServer()
// }
//
// 定义helloService并实现约定的接口
type helloService struct {
	gw.UnimplementedHelloHTTPServer
}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *gw.HelloHTTPRequest) (*gw.HelloHTTPResponse, error) {
	resp := new(gw.HelloHTTPResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func grpc_server() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	gw.RegisterHelloHTTPServer(s, HelloService)

	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}

func main() {
	go grpc_server()

	// 1. 定义一个context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// grpc服务地址
	endpoint := "127.0.0.1:50052"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP转grpc
	err := gw.RegisterHelloHTTPHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		grpclog.Fatalf("Register handler err:%v\n", err)
	}

	fmt.Println("HTTP Listen on 8080")
	http.ListenAndServe(":8080", mux)
}
