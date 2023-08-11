package port

import (
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
)

type UserServiceInterface interface {
	Register(user domain.UserEntity) (*web.UserResponseSuccess, error)
	Update(id uint, userUpt domain.UserEntity) (*web.UserResponseSuccess, error)
	FindByEmail(email string) (*web.UserResponseSuccess, error)
	DeleteAccount(id uint) error
	JoinEvent(join domain.ParticipantEntity) (*web.JoinEventResponseSuccess, error)
	DetailUser(id uint) (*web.DetailUserResponse, error)
}