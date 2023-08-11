package services

import (
	"context"

	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
	"github.com/rulanugrh/orion/repository/port"
	portServ "github.com/rulanugrh/orion/services/port"
)

type eventservices struct {
	eventRepo port.EventRepositoryInterface
}

func NewEventServices(event port.EventRepositoryInterface) portServ.EventServiceInterface {
	return &eventservices{
		eventRepo: event,
	}
}

func (srv *eventservices) CreateEvent(event domain.EventEntity) (*web.EventResponseSuccess, error) {
	result, err := srv.eventRepo.CreateEvent(context.Background(), event)
	if err != nil {
		return nil, err
	}
	var participantList []web.ParcipantList
	for _, data := range result.Participant {
		participant := web.ParcipantList{
			UserName: data.User.Name,
			Email:    data.User.Email,
		}

		participantList = append(participantList, participant)
	}

	response := web.EventResponseSuccess{
		Name:        result.Name,
		Description: result.Description,
		Owner:       result.User.Name,
		CreateAt:    result.CreatedAt,
		UpdateAt:    result.UpdatedAt,
		Parcipant:   participantList,
	}

	return &response, nil
}

func (srv *eventservices) GetEventByID(id uint) (*web.EventResponseSuccess, error) {
	result, err := srv.eventRepo.GetEventById(context.Background(), id)
	if err != nil {
		return nil, err
	}

	var commentList []web.CommentList
	for _, commentRes := range result.Comments {
		comment := web.CommentList{
			UserName: commentRes.User.Name,
			Comment:  commentRes.Comment,
		}

		commentList = append(commentList, comment)
	}

	var participantList []web.ParcipantList
	for _, data := range result.Participant {
		participant := web.ParcipantList{
			UserName: data.User.Name,
			Email:    data.User.Email,
		}

		participantList = append(participantList, participant)
	}

	response := web.EventResponseSuccess{
		Name:        result.Name,
		Description: result.Description,
		Owner:       result.User.Name,
		CreateAt:    result.CreatedAt,
		UpdateAt:    result.UpdatedAt,
		Comment:     commentList,
		Parcipant:   participantList,
	}

	return &response, nil
}

func (srv *eventservices) GetEvent() ([]web.EventResponseSuccess, error) {
	result, err := srv.eventRepo.GetEvent(context.Background())
	if err != nil {
		return []web.EventResponseSuccess{}, nil
	}

	var response []web.EventResponseSuccess
	for _, data := range result {

		var commentList []web.CommentList
		for _, comment := range data.Comments {
			comments := web.CommentList{
				UserName: comment.User.Name,
				Comment:  comment.Comment,
			}

			commentList = append(commentList, comments)
		}
		var participantList []web.ParcipantList
		for _, res := range data.Participant {
			participant := web.ParcipantList{
				UserName: res.User.Name,
				Email:    res.User.Email,
			}

			participantList = append(participantList, participant)
		}

		res := web.EventResponseSuccess{
			Name:        data.Name,
			Description: data.Description,
			Owner:       data.User.Name,
			CreateAt:    data.CreatedAt,
			UpdateAt:    data.UpdatedAt,
			Comment:     commentList,
			Parcipant:   participantList,
		}

		response = append(response, res)
	}

	return response, nil
}

func (srv *eventservices) UpdateEvent(id uint, eventUpt domain.EventEntity) (*web.EventResponseSuccess, error) {
	result, err := srv.eventRepo.UpdateEvent(context.Background(), id, eventUpt)
	if err != nil {
		return nil, err
	}

	var commentList []web.CommentList
	for _, data := range result.Comments {
		comment := web.CommentList{
			UserName: data.User.Name,
			Comment:  data.Comment,
		}

		commentList = append(commentList, comment)
	}
	var participantList []web.ParcipantList
	for _, data := range result.Participant {
		participant := web.ParcipantList{
			UserName: data.User.Name,
			Email:    data.User.Email,
		}

		participantList = append(participantList, participant)
	}

	response := web.EventResponseSuccess{
		Name:        result.Name,
		Description: result.Description,
		Owner:       result.User.Name,
		CreateAt:    result.CreatedAt,
		UpdateAt:    result.UpdatedAt,
		Comment:     commentList,
		Parcipant:   participantList,
	}

	return &response, nil

}

func (srv *eventservices) DeleteEvent(id uint) error {
	err := srv.eventRepo.DeleteEvent(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}
