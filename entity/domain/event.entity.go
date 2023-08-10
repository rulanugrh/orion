package domain

import "gorm.io/gorm"

type EventEntity struct {
	gorm.Model
	Name 			string 			`json:"name" form:"name" validate:"required"`
	Description 	string 			`json:"desc" form:"desc" validate:"required"`
	CreatorID 		uint 			`json:"user_id" form:"user_id" validate:"required"`
	User			UserEntity 		`json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Participant  	[]UserEntity	`json:"parcipant" form:"participant" gorm:"many2many:participant_event"`
	Comments		[]CommentEntity `json:"comment" form:"comment" gorm:"many2many:comment_event"`
}