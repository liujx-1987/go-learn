# 编译Product.proto之后输出到service文件夹
protoc --go_out=../service Product.proto

protoc --go_out=plugins=grpc:../service Product.proto