package grpc

import (
	"github.com/ngdangkietse/ndk-go-server/grpc/servers/account"
	"google.golang.org/grpc"
)

func NewGrpcServer(server *grpc.Server) {
	account.NewGrpcHandler().RegisterAccountGrpcServer(server)
}
