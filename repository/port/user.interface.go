package port

import (
	"context"

	"github.com/rulanugrh/orion/entity/domain"
)

type UserRepositoryInterface interface {
	Register(ctx context.Context, user domain.UserEntity) (*domain.UserEntity, error)
	Update(ctx context.Context, id uint, userUpt domain.UserEntity) (*domain.UserEntity, error)
	FindByEmail(ctx context.Context, email string) (*domain.UserEntity, error)
	DeleteAccount(ctx context.Context, id uint) error
}