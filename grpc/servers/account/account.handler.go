package account

import (
	"github.com/ngdangkietse/ndk-go-proto/generated/account"
	implRepo "github.com/ngdangkietse/ndk-go-server/data/repositories/impl"
	"github.com/ngdangkietse/ndk-go-server/grpc/services/impl"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
}

func NewGrpcHandler() *GrpcHandler {
	return &GrpcHandler{}
}

func (handler *GrpcHandler) RegisterAccountGrpcServer(server *grpc.Server) {
	userRepository := implRepo.NewUserRepository()
	userService := impl.NewUserService(userRepository)
	account.RegisterUserServiceServer(server, NewAccountGrpcServer(userService))
}
