package web

import (
	"time"
)

type EventResponseSuccess struct {
	Name        string 			`json:"name"`
	Description string 			`json:"description"`
	Owner       string 			`json:"owner"`
	CreateAt    time.Time 		`json:"create_at"`
	UpdateAt	time.Time 		`json:"update_at"`
	Comment 	[]CommentList 	`json:"comment"`
	Parcipant 	[]ParcipantList `json:"participant"`
}

type JoinEventResponseSuccess struct {
	EventName string `json:"event_name"`
	EventDesc string `json:"event_desc"`
	UserName string `json:"user_name"`
}

type UserResponseSuccess struct {
	Name 	string `json:"name"`
	Message string `json:"message"`
}

type DetailUserResponse struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Notelp string `json:"no_telepon"`
	Email string `json:"email"`
	Event []EventList `json:"event_joining"`
}

type EventList struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type CommentResponseSuccess struct {
	EventName string `json:"event_name"`
	EventDesc string `json:"event_description"`
	UserName  string `json:"user_name"`
	Comment   string `json:"comment"`
}

type CommentList struct {
	UserName string 	`json:"user_name"`
	Comment  string 	`json:"comment"`
}

type ParcipantList struct {
	UserName string `json:"user_name"`
	Email string `json:"user_email"`
}

type ResponseFailure struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}