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
	"github.com/rulanugrh/orion/services/port"
)

type eventhandler struct {
	eventService port.EventServiceInterface
}

func NewEventHandler(event port.EventServiceInterface) portHand.EventHandlerInterface {
	return &eventhandler{
		eventService: event,
	}
}
func (hnd *eventhandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var req domain.EventEntity
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	result, err := hnd.eventService.CreateEvent(req)
	if err != nil {
		res := web.WebValidationError{
			Message: "Cant Create Event",
			Errors:  err,
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Create Event Because %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	} else {
		res := web.ResponseSuccess{
			Message: "Success create event",
			Data:    result,
		}
		response, _ := json.Marshal(res)

		w.Write(response)
	}

}

func (hnd *eventhandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	getId := mux.Vars(r)
	paramsId := getId["id"]
	id, _ := strconv.Atoi(paramsId)

	result, err := hnd.eventService.GetEventByID(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Find Event by this ID",
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Find Event by this ID %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	} else {
		res := web.ResponseSuccess{
			Message: "Success find event",
			Data:    result,
		}
		response, _ := json.Marshal(res)

		w.Write(response)
	}

}

func (hnd *eventhandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	result, err := hnd.eventService.GetEvent()
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Find Event",
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Find Event %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	} else {
		res := web.ResponseSuccess{
			Message: "Success find event",
			Data:    result,
		}
		response, _ := json.Marshal(res)

		w.Write(response)
	}

}

func (hnd *eventhandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var req domain.EventEntity
	data, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(data, &req)
	getId := mux.Vars(r)
	paramsId := getId["id"]
	id, _ := strconv.Atoi(paramsId)

	result, err := hnd.eventService.UpdateEvent(uint(id), req)
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Update Event",
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Update Event, because %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	} else {
		res := web.ResponseSuccess{
			Message: "Success update event",
			Data:    result,
		}
		response, _ := json.Marshal(res)

		w.Write(response)
	}

}

func (hnd *eventhandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	getId := mux.Vars(r)
	paramsId := getId["id"]
	id, _ := strconv.Atoi(paramsId)

	err := hnd.eventService.DeleteEvent(uint(id))
	if err != nil {
		res := web.ResponseFailure{
			Message: "Cant Delete Event by this ID",
		}

		response, _ := json.Marshal(res)
		log.Printf("Cant Delete Event by this ID, because %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	} else {
		res := web.ResponseSuccess{
			Message: "Success delete event",
			Data:    "Data is has been delete",
		}
		response, _ := json.Marshal(res)

		w.Write(response)
	}

}
