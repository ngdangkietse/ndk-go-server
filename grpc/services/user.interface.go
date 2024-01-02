package services

import (
	"context"
	"github.com/ngdangkietse/ndk-go-proto/generated/account"
	"github.com/ngdangkietse/ndk-go-proto/generated/common"
)

type IUserService interface {
	UpsertUser(ctx context.Context, req *account.PUser) (*account.PUpsertUserResponse, error)
	FindUserById(ctx context.Context, req *common.PIdRequest) (*account.PGetUserResponse, error)
	FindUserByEmail(ctx context.Context, req *common.PEmailRequest) (*account.PGetUserResponse, error)
}
