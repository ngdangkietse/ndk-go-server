package account

import (
	"context"
	"github.com/ngdangkietse/ndk-go-proto/generated/account"
	"github.com/ngdangkietse/ndk-go-server/grpc/services"
)

type GrpcServer struct {
	account.UnimplementedUserServiceServer
	userService services.IUserService
}

func NewAccountGrpcServer(userService services.IUserService) *GrpcServer {
	return &GrpcServer{userService: userService}
}

func (gs *GrpcServer) UpsertUser(ctx context.Context, req *account.PUser) (*account.PUpsertUserResponse, error) {
	return gs.userService.UpsertUser(ctx, req)
}
