package src

//go gRPC tutorial
// 1 Define a service in a .proto file
// 2 generate server and client code using the protocal buffer compiler
// 3 use the go gRPC API to write a simple client and server for you service

// protocal buffers
// proto3 \

// why use gRPC

// install gRPC
// go get -u google.golang.org/grpc
// 国外的 vps 只有windows　的openvpn
// 把源码下载到 服务器　gopath/src ,再用tar压缩.　在用　scp 下载到linux机器上
// scp 制定端口 使用大写P
//scp -P 27196 xxxx@xxx.xxx.xxx.xxx:/root/go/src.tar ./

// 接下来两步没安装　等到用时有问题再看
// https://github.com/google/protobuf/releases
// Unzip this file.
// Update the environment variable PATH to include the path to the protoc binary file.
//  go get -u github.com/golang/protobuf/protoc-gen-go
// 打开两个　terminal
// cd $GOPATH/src/google.golang.org/grpc/examples/helloworld
// go run greeter_server/main.go
// go run greeter_client/main.go

// 转移阵地。　转到 microservices