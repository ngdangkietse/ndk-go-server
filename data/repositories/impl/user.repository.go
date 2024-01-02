package impl

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietse/ndk-go-server/data/models"
	"github.com/ngdangkietse/ndk-go-server/data/repositories"
	"github.com/ngdangkietse/ndk-go-server/db"
	"gorm.io/gorm"
)

type UserRepository struct {
	instance *gorm.DB
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepository{
		instance: db.Instance,
	}
}

func (u UserRepository) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	record := u.instance.Where("id = ?", id).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}
	return &user, nil
}

func (u UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	record := u.instance.Where("email = ?", email).First(&user)
	if record.Error != nil {
		return nil, record.Error
	}
	return &user, nil
}
