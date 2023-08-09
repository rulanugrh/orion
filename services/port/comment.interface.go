package port

import (
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
)

type CommentServiceInterface interface {
	CreateComment(comment domain.CommentEntity) (*web.CommentResponseSuccess, error)
	GetAllComment() ([]web.CommentResponseSuccess, error)
}