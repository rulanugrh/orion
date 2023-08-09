package services

import (
	"context"

	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
	"github.com/rulanugrh/orion/repository/port"
	portServ "github.com/rulanugrh/orion/services/port"
)
type userservices struct {
	userRepo port.UserRepositoryInterface
}

func NewUserServices(user port.UserRepositoryInterface) portServ.UserServiceInterface {
	return &userservices{
		userRepo: user,
	}
}

func (srv *userservices) Register(user domain.UserEntity) (*web.UserResponseSuccess, error) {
	result, err := srv.userRepo.Login(context.Background(), user)
	if err != nil {
		return nil, err
	}

	response := web.UserResponseSuccess {
		Name: result.Name,
		Message: "Success Register",
	}

	return &response, nil
}

func (srv *userservices) Update(id uint, userUpt domain.UserEntity) (*web.UserResponseSuccess, error) {
	result, err := srv.userRepo.Update(context.Background(), id, userUpt)
	if err != nil {
		return nil, err
	}

	response := web.UserResponseSuccess {
		Name: result.Name,
		Message: "Success Update",
	}

	return &response, nil
}

func (srv *userservices) FindByEmail(email string) (*web.UserResponseSuccess, error) {
	result, err := srv.userRepo.FindByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	response := web.UserResponseSuccess {
		Name: result.Name,
		Message: "Success Login",
	}

	return &response, nil
}

func (srv *userservices) DeleteAccount(id uint) error {
	if err := srv.userRepo.DeleteAccount(context.Background(), id); err != nil {
		return err
	}
	return nil
}