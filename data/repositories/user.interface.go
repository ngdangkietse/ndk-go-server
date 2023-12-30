package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietse/ndk-go-server/data/models"
)

type IUserRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
