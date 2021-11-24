#! /bin/bash

//先安装环境
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc




# 编译Product.proto之后输出到service文件夹
#protoc --go_out=../service Product.proto

# 编译Product.proto之后输出到Product文件夹
protoc --go_out=plugins=grpc:./Product Product/Product.proto

//hello_http工程
$ cd proto

# 编译google.api
$ protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

# 编译hello_http.proto

$ protoc -I . \
    --go_out ./ --go_opt paths=source_relative \
    --go-grpc_out ./ --go-grpc_opt paths=source_relative \
    hello_http/*.proto
 

# 编译hello_http.proto gateway

$ protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    hello_http/*.proto



