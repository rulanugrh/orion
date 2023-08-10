package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
	portHand "github.com/rulanugrh/orion/handler/port"
	"github.com/rulanugrh/orion/services/port"
)

type commenthandler struct {
	commentService port.CommentServiceInterface
}

func NewCommentHandler(comment port.CommentServiceInterface) portHand.CommentHandlerInterface {
	return &commenthandler{
		commentService: comment,
	}
}

func (hnd *commenthandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var req domain.CommentEntity
	data, _ := ioutil.ReadAll(r.Body)
	
	json.Unmarshal(data, &req)

	result, err := hnd.commentService.CreateComment(req)
	if err != nil {
		res := web.ResponseFailure {
			Message: "Cant Create Comment",
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Create Comment Because %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess {
		Message: "Success create comment",
		Data: result,
	}
	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (hnd *commenthandler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	result, err := hnd.commentService.GetAllComment()
	if err != nil {
		res := web.ResponseFailure {
			Message: "Cant Get All Comment",
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Get All Comment Because %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess {
		Message: "Success get all comment",
		Data: result,
	}
	response, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}