package port

import (
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
)
type EventServiceInterface interface {
	CreateEvent(event domain.EventEntity) (*web.EventResponseSuccess, error)
	GetEvent() ([]web.EventResponseSuccess, error)
	GetEventByID(id uint) (*web.EventResponseSuccess, error)
	UpdateEvent(id uint, event domain.EventEntity) (*web.EventResponseSuccess, error)
	DeleteEvent(id uint) error
}