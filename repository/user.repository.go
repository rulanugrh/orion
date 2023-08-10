package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/orion/configs"
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/repository/port"
)

type userrepository struct{}

func NewUserRepository() port.UserRepositoryInterface {
	return &userrepository{}
}

func (rep *userrepository) Register(ctx context.Context, user domain.UserEntity) (*domain.UserEntity, error) {
	err := configs.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return &user, nil
}

func (rep *userrepository) Update(ctx context.Context, id uint, userUpt domain.UserEntity) (*domain.UserEntity, error) {
	var user domain.UserEntity
	err := configs.DB.WithContext(ctx).Model(&userUpt).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return &user, nil
}

func (rep *userrepository) FindByEmail(ctx context.Context, email string) (*domain.UserEntity, error) {
	var user domain.UserEntity
	err := configs.DB.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return &user, nil
}

func (rep *userrepository) DeleteAccount(ctx context.Context, id uint) error {
	var user domain.UserEntity
	err := configs.DB.WithContext(ctx).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return nil
}

func (rep *userrepository) JoinEvent(ctx context.Context, id uint) error {
	var user domain.UserEntity
	errs := configs.DB.WithContext(ctx).Model(&user.Event).Where("id = ?", id).Association("Participant").Append(&user)
	if errs != nil {
		log.Printf("Found Error %v", errs)
	}

	return nil
}