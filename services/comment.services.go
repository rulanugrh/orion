package services

import (
	"context"

	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
	"github.com/rulanugrh/orion/middleware"
	"github.com/rulanugrh/orion/repository/port"
	portServ "github.com/rulanugrh/orion/services/port"
)

type commentservice struct {
	commentRepo port.CommentRepositoryInterface
}

func NewCommentService(comment port.CommentRepositoryInterface) portServ.CommentServiceInterface {
	return &commentservice{
		commentRepo: comment,
	}
}

func (srv *commentservice) CreateComment(comment domain.CommentEntity) (*web.CommentResponseSuccess, error) {
	errStruct := middleware.ValidateStruct(comment)
	if errStruct != nil {
		return nil, errStruct
	}

	result, err := srv.commentRepo.CreateComment(context.Background(), comment)
	if err != nil {
		return nil, err
	}

	response := web.CommentResponseSuccess{
		EventName: result.Event.Name,
		EventDesc: result.Event.Description,
		UserName:  result.User.Name,
		Comment:   result.Comment,
	}

	return &response, nil
}

func (srv *commentservice) GetAllComment() ([]web.CommentResponseSuccess, error) {
	result, err := srv.commentRepo.GetAllComment(context.Background())
	if err != nil {
		return []web.CommentResponseSuccess{}, nil
	}

	var response []web.CommentResponseSuccess
	for _, data := range result {
		comment := web.CommentResponseSuccess{
			EventName: data.Event.Name,
			EventDesc: data.Event.Description,
			UserName:  data.User.Name,
			Comment:   data.Comment,
		}

		response = append(response, comment)
	}

	return response, nil
}
