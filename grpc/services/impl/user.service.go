package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ngdangkietse/ndk-go-proto/generated/account"
	"github.com/ngdangkietse/ndk-go-proto/generated/common"
	"github.com/ngdangkietse/ndk-go-server/data/repositories"
	"github.com/ngdangkietse/ndk-go-server/grpc/services"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) services.IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) UpsertUser(ctx context.Context, req *account.PUser) (*account.PUpsertUserResponse, error) {
	return &account.PUpsertUserResponse{
		Code:    common.PCode_SUCCESS,
		Message: fmt.Sprintf("Create/Update user [%v] successful!", req.Email),
	}, nil
}

func (u *UserService) FindUserById(ctx context.Context, req *common.PIdRequest) (*account.PUser, error) {
	//TODO implement me
	data, err := u.userRepository.FindById(ctx, uuid.MustParse(req.Id))
	if err != nil {
		return nil, err
	}
	return &account.PUser{
		Id:       data.Id.String(),
		FullName: data.FullName,
		Email:    data.Email,
	}, nil
}

func (u *UserService) FindUserByEmail(ctx context.Context, req *common.PEmailRequest) (*account.PUser, error) {
	//TODO implement me
	data, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &account.PUser{
		Id:       data.Id.String(),
		FullName: data.FullName,
		Email:    data.Email,
	}, nil
}
