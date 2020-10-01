package main

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc"
	"github.com/spf13/viper"
)

// Generate gRPC code from protobuf files
//go:generate protoc -I. -Igrpc/third_party/googleapis --go_out . --go_opt plugins=grpc --go_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --swagger_out . --swagger_opt logtostderr=true --grpc-gateway_opt paths=source_relative grpc/proto/urls/urls.proto

func main() {
	viper.SetDefault("port", 8080)
	viper.SetDefault("grpcPort", 8081)

	//rest.RunRESTServer()
	grpc.RunGRPCServer()
}
