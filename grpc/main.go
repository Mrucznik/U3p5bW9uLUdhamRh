package main

import (
	"context"
	"fmt"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/proto/urls"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

//go:generate protoc -I. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_opt plugins=grpc --grpc-gateway_out . --go_out . --grpc-gateway_opt logtostderr=true --swagger_out ./gen/swagger --swagger_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/urls/urls.proto

func main() {
	viper.SetDefault("port", 8080)
	viper.SetDefault("grpcPort", 8081)

	// if we crash the go code, we get the file name and line number in log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logrus.Infoln("Starting listener for gRPC API")
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("grpcPort")))
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	logrus.Infoln("Listener started on", listener.Addr())

	//server options
	var opts []grpc.ServerOption
	if viper.GetBool("ssl.enabled") {
		creds, sslErr := credentials.NewServerTLSFromFile(
			viper.GetString("ssl.cert"), viper.GetString("ssl.key"))
		if sslErr != nil {
			logrus.Fatalln("Failed loading certificates:", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	//serve gRPC services
	s := grpc.NewServer(opts...)
	reflection.Register(s)
	defer s.Stop()

	// Register services
	urls.RegisterUrlsServiceServer(s, service.NewServer())

	go func() {
		logrus.Println("Starting server.")

		if err := s.Serve(listener); err != nil {
			logrus.Fatalln("Failed to serve", err)
		}
	}()

	// Set up the gRPC gateway for REST endpoints
	if err = setUpgRPCGateway(); err != nil {
		logrus.Fatalln("Failed to set up gRPC gateway: ", err)
	}

	// Wait for CTRL+C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch // Block until signal is received
	logrus.Infoln("\nStopping the server.")
}

func setUpgRPCGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := urls.RegisterUrlsServiceHandlerFromEndpoint(ctx, mux,
		fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("grpcPort")),
		opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port")), mux)
}
