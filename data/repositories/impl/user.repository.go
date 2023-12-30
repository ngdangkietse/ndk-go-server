package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietse/ndk-go-server/data/models"
	"github.com/ngdangkietse/ndk-go-server/data/repositories"
)

type UserRepository struct {
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepository{}
}

func (u UserRepository) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}
