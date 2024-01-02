package main

import (
	"github.com/ngdangkietse/ndk-go-server/config"
	"github.com/ngdangkietse/ndk-go-server/db"
	server "github.com/ngdangkietse/ndk-go-server/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.InitConfig()
	db.ConnectDB(config.EnvConfig)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server.NewGrpcServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
