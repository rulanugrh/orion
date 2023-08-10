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

type UserResponseSuccess struct {
	Name 	string `json:"name"`
	Message string `json:"message"`
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