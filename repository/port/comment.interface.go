package port

import (
	"context"

	"github.com/rulanugrh/orion/entity/domain"
)

type CommentRepositoryInterface interface {
	CreateComment(ctx context.Context, comment domain.CommentEntity) (*domain.CommentEntity, error)
	GetAllComment(ctx context.Context) ([]domain.CommentEntity, error)
}