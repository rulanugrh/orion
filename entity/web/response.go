package web

import (
	"time"
)

type EventResponseSuccess struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Owner       string          `json:"owner"`
	CreateAt    time.Time       `json:"create_at"`
	UpdateAt    time.Time       `json:"update_at"`
	Comment     []CommentList   `json:"comment"`
	Parcipant   []ParcipantList `json:"participant"`
}

type ValidationList struct {
	Field string      `json:"field"`
	Error interface{} `json:"error"`
}

type ValidationError struct {
	Message string           `json:"message"`
	Errors  []ValidationList `json:"error"`
}

func (err ValidationError) Error() string {
	return err.Message
}

type WebValidationError struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"error"`
}

type JoinEventResponseSuccess struct {
	EventName string `json:"event_name"`
	EventDesc string `json:"event_desc"`
	UserName  string `json:"user_name"`
}

type UserResponseSuccess struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type DetailUserResponse struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Notelp  string `json:"no_telepon"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Avatar  string `json:"avatar"`
}

type CommentResponseSuccess struct {
	EventName string `json:"event_name"`
	EventDesc string `json:"event_description"`
	UserName  string `json:"user_name"`
	Comment   string `json:"comment"`
}

type CommentList struct {
	UserName string `json:"user_name"`
	Comment  string `json:"comment"`
}

type ParcipantList struct {
	UserName string `json:"user_name"`
	Email    string `json:"user_email"`
}

type ResponseFailure struct {
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
