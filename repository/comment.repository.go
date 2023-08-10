package repository

import (
	"context"
	"log"

	"github.com/rulanugrh/orion/configs"
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/repository/port"
)

type commentrepositoy struct{}

func NewCommentRepository() port.CommentRepositoryInterface {
	return &commentrepositoy{}
}

func (rep *commentrepositoy) CreateComment(ctx context.Context, comment domain.CommentEntity) (*domain.CommentEntity, error) {
	err := configs.DB.WithContext(ctx).Create(&comment).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	errs := configs.DB.WithContext(ctx).Preload("Event").Preload("User").Find(&comment).Error
	if errs != nil {
		log.Printf("Found Error %v", errs)
	}

	errsEvent := configs.DB.Model(&comment.Event).Association("Comments").Append(&comment)
	if errsEvent != nil {
		log.Printf("Found Error %v", errs)
	}
	
	return &comment, nil
}

func (rep *commentrepositoy) GetAllComment(ctx context.Context) ([]domain.CommentEntity, error) {
	var comment []domain.CommentEntity
	err := configs.DB.WithContext(ctx).Preload("Event").Preload("User").Find(&comment).Error
	if err != nil {
		log.Printf("Found Error %v", err)
	}

	return comment, nil
}