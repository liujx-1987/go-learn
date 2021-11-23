package main

import (
	"CoolCar/proto/product"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.SetFlags(log.Lshortfile)
	// 1. 添加公钥证书的引用， codepie.fun是之前生成证书的时候填写的common name
	tls, err := credentials.NewClientTLSFromFile("../../keys/cert.pem", "")
	if err != nil {
		fmt.Println("客户端获取证书失败")
		log.Fatal("客户端获取证书失败: ", err)
	}

	// 2. 新建连接，端口是服务端开放的8082端口
	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(tls))
	if err != nil {
		log.Fatal(err)
	}

	// 退出时关闭链接
	defer conn.Close()

	// 3. 调用Product.pb.go中的NewProdServiceClient方法
	productServiceClient := product.NewProdServiceClient(conn)

	// 4. 直接像调用本地方法一样调用GetProductStock方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := productServiceClient.GetProductStock(ctx, &product.ProductRequest{ProdId: 233})
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("调用gRPC方法成功，ProdStock = ", resp.ProdStock)
	fmt.Println("client")
}
