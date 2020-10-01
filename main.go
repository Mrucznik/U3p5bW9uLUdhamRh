package main

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc"
	"github.com/spf13/viper"
)

// Generate gRPC code from protobuf files
//go:generate protoc -I. -Igrpc/third_party/googleapis --go_out . --go_opt plugins=grpc --go_opt paths=source_relative --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --swagger_out . --swagger_opt logtostderr=true --grpc-gateway_opt paths=source_relative grpc/proto/urls/urls.proto

func main() {
	viper.SetDefault("DSN", "root:secret@tcp(mysqld:3306)/mysql")
	viper.SetDefault("USE_DATABASE", true)
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("GRPC_PORT", 8081)

	//rest.RunRESTServer()
	grpc.RunGRPCServer()
}
