package main

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/rest"
)

func main() {
	rest.RunRESTServer()
	grpc.RunGRPCServer()
}
