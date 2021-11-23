package main

import (
	// "gomicro-quickstart/grpc_demo/service"
	"fmt"
	"log"
	"net"

	"CoolCar/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 1. 引用证书
	tls, err := credentials.NewServerTLSFromFile("./keys/cert.pem", "./keys/key.pem")
	if err != nil {
		log.Fatal("服务端获取证书失败: ", err)
	}

	// 2. new一个grpc的server，并且加入证书
	rpcServer := grpc.NewServer(grpc.Creds(tls))

	// 3. 将刚刚我们新建的ProdService注册进去
	service.RegisterProdServiceServer(rpcServer, new(service.ProductService))

	// 4. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	fmt.Println("start grpc Server")
	// 5. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)

}
