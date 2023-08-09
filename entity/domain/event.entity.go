package domain

import "gorm.io/gorm"

type EventEntity struct {
	gorm.Model
	Name 			string 			`json:"name" form:"name" validate:"required"`
	Description 	string 			`json:"desc" form:"desc" validate:"required"`
	UserID 			uint 			`json:"user_id" form:"user_id" validate:"required"`
	User			UserEntity 		`json:"user" form:"user"`
	Comments		[]CommentEntity `json:"comment" form:"comment"`
}