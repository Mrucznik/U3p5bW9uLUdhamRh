package grpc

import (
	"context"
	"fmt"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/db"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine/database"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine/in_memory"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func RunGRPCServer() {
	// if we crash the go code, we get the file name and line number in log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logrus.Infoln("Starting listener for gRPC API")
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("GRPC_PORT")))
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	logrus.Infoln("Listener started on", listener.Addr())

	//serve gRPC services
	s := grpc.NewServer(grpc.MaxRecvMsgSize(1048576))
	reflection.Register(s)
	defer s.Stop()

	// Register services
	if viper.GetBool("USE_DATABASE") {
		logrus.Infoln("Connecting to MySQL databasee.")
		mysqlConnection := db.SetUpDatabase()
		defer mysqlConnection.Close()
		logrus.Infoln("Connectd.")

		urls.RegisterUrlsServiceServer(s, NewServer(database.NewURLsService(mysqlConnection)))
	} else {
		urls.RegisterUrlsServiceServer(s, NewServer(in_memory.NewURLsService()))
	}

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
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := urls.RegisterUrlsServiceHandlerFromEndpoint(ctx, mux,
		fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("GRPC_PORT")),
		opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", viper.GetString("HOST"), viper.GetInt("PORT")), mux)
}
