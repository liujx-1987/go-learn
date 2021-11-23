#! /bin/bash


# 编译Product.proto之后输出到service文件夹
#protoc --go_out=../service Product.proto

# 编译Product.proto之后输出到Product文件夹
protoc --go_out=plugins=grpc:./Product Product/Product.proto