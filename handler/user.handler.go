package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/orion/entity/domain"
	"github.com/rulanugrh/orion/entity/web"
	portHand "github.com/rulanugrh/orion/handler/port"
	"github.com/rulanugrh/orion/middleware"
	"github.com/rulanugrh/orion/services/port"
)

type userhandler struct {
	userServ port.UserServiceInterface
}

func NewUserHandler(user port.UserServiceInterface) portHand.UserHandlerInterface {
	return &userhandler{
		userServ: user,
	}
}
func (hnd *userhandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.UserEntity
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	password := middleware.HashPassword(req.Password)
	req.Password = password

	_, err := hnd.userServ.Register(req)
	if err != nil {
		res := web.WebValidationError{
			Message: "Cant create account or Email has been used",
			Errors:  err,
		}
		log.Printf("cant create because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

	} else {
		token, errToken := middleware.GenerateToken(req)
		if errToken != nil {
			res := web.ResponseFailure{
				Message: "Cant Generate Token",
			}
			log.Printf("cant generate token because: %v", errToken)
			response, _ := json.Marshal(res)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
		} else {
			res := web.ResponseSuccess{
				Message: "success register account",
				Data:    token,
			}
			response, _ := json.Marshal(res)
			w.Write(response)
		}
	}

}

func (hnd *userhandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req domain.UserEntity
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	_, err := hnd.userServ.FindByEmail(req.Email)
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Find Account with this Email",
		}
		log.Printf("cant find account because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
	token, errToken := middleware.GenerateToken(req)
	if errToken != nil {
		res := web.ResponseFailure{
			Message: "Cant Generate Token",
		}
		log.Printf("cant generate token because: %v", errToken)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Message: "success refreh token ",
		Data:    token,
	}
	response, _ := json.Marshal(res)

	w.Write(response)
}

func (hnd *userhandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserEntity
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	result, err := hnd.userServ.FindByEmail(req.Email)
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Find Account with this Email",
		}
		log.Printf("cant find account because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	compare := []byte(req.Password)
	if errCheck := middleware.ComparePassword(req.Password, compare); errCheck != nil {
		res := web.ResponseFailure{
			Message: "Password not matched",
		}
		log.Printf("password not matched because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Message: "success login account",
		Data:    result,
	}
	response, _ := json.Marshal(res)

	w.Write(response)

}

func (hnd *userhandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	getId := mux.Vars(r)
	params := getId["id"]
	id, _ := strconv.Atoi(params)

	var req domain.UserEntity
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &req)

	result, err := hnd.userServ.Update(uint(id), req)
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Update Account with this ID",
		}
		log.Printf("cant update account because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Message: "success update account",
		Data:    result,
	}
	response, _ := json.Marshal(res)

	w.Write(response)
}

func (hnd *userhandler) JoinEvent(w http.ResponseWriter, r *http.Request) {
	var req domain.ParticipantEntity
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &req)

	result, err := hnd.userServ.JoinEvent(req)
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Join Event with This ID",
		}
		log.Printf("cant join event because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Message: "success join event",
		Data:    result,
	}
	response, _ := json.Marshal(res)

	w.Write(response)
}

func (hnd *userhandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	getId := mux.Vars(r)
	params := getId["id"]
	id, _ := strconv.Atoi(params)

	if err := hnd.userServ.DeleteAccount(uint(id)); err != nil {
		res := web.ResponseFailure{
			Message: "Cant Delete Account with this ID",
		}
		log.Printf("cant delete account because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Message: "success delete account",
		Data:    "Account with this ID has been delete",
	}
	response, _ := json.Marshal(res)

	w.Write(response)
}

func (hnd *userhandler) DetailUser(w http.ResponseWriter, r *http.Request) {
	getId := mux.Vars(r)
	params := getId["id"]
	id, _ := strconv.Atoi(params)

	result, err := hnd.userServ.DetailUser(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant see detail with This ID",
		}
		log.Printf("cant see detail because: %v", err)
		response, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}

	res := web.ResponseSuccess{
		Message: "success get detail",
		Data:    result,
	}
	response, _ := json.Marshal(res)

	w.Write(response)
}
