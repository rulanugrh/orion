package port

import (
	"context"

	"github.com/rulanugrh/orion/entity/domain"
)

type EventRepositoryInterface interface {
	CreateEvent(ctx context.Context, event domain.EventEntity) (*domain.EventEntity, error)
	GetEventById(ctx context.Context, id uint) (*domain.EventEntity, error)
	GetEvent(ctx context.Context) ([]domain.EventEntity, error)
	UpdateEvent(ctx context.Context, id uint, eventUpt domain.EventEntity) (*domain.EventEntity, error)
	DeleteEvent(ctx context.Context, id uint) error
}