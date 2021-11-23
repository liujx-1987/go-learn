package main

import (
	// "gomicro-quickstart/grpc_demo/service"
	"fmt"
	"log"
	"net"
	"net/http"

	"CoolCar/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// grpc_main()
	grpc_http()

}

func grpc_main() {
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

func grpc_http() {
	// 1. 引用证书
	tls, err := credentials.NewServerTLSFromFile("./keys/cert.pem", "./keys/key.pem")
	if err != nil {
		log.Fatal("服务端获取证书失败: ", err)
	}

	// 2. new一个grpc的server，并且加入证书
	rpcServer := grpc.NewServer(grpc.Creds(tls))

	// 3. 将刚刚我们新建的ProdService注册进去
	service.RegisterProdServiceServer(rpcServer, new(service.ProductService))

	// 4. 新建一个路由，并传入rpcServer
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})

	// 5. 定义httpServer，监听8082
	httpServer := http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	fmt.Println("start grpc Server")

	// 6. 以https形式监听httpServer
	httpServer.ListenAndServeTLS("./keys/cert.pem", "./keys/key.pem")

}
