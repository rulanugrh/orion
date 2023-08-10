package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/orion/configs"
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/repository/port"
)
type eventrepository  struct {}

func NewEventRepository() port.EventRepositoryInterface {
	return &eventrepository{}
}

func (rep *eventrepository) CreateEvent(ctx context.Context, event domain.EventEntity) (*domain.EventEntity, error) {
	err := configs.DB.WithContext(ctx).Create(&event).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	errs := configs.DB.WithContext(ctx).Preload("User").Find(&event).Error
	if errs != nil {
		log.Printf("Found Error %v", errs)
	}
	
	return &event, nil
}



func (rep *eventrepository) GetEventById(ctx context.Context, id uint) (*domain.EventEntity, error) {
	var event domain.EventEntity
	err := configs.DB.WithContext(ctx).Preload("User").Preload("Comments").	Where("id = ?", id).Find(&event).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return &event, nil
}

func (rep *eventrepository) GetEvent(ctx context.Context) ([]domain.EventEntity, error) {
	var events []domain.EventEntity

	err := configs.DB.WithContext(ctx).Preload("User").Preload("Comments").Preload("Comments.User").Find(&events).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return events, nil
}

func (rep *eventrepository) UpdateEvent(ctx context.Context, id uint, eventUpt domain.EventEntity) (*domain.EventEntity, error) {
	var event domain.EventEntity

	err := configs.DB.WithContext(ctx).Model(&eventUpt).Where("id = ?", id).Updates(&event).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return &event, nil
}

func (rep *eventrepository) DeleteEvent(ctx context.Context, id uint) error {
	var event domain.EventEntity
	err := configs.DB.WithContext(ctx).Where("id = ?", id).Delete(&event).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return nil
}