package domain

import "gorm.io/gorm"

type CommentEntity struct {
	gorm.Model
	UserID  uint        `json:"user_id" form:"user_id" validate:"required"`
	EventID uint        `json:"event_id" form:"event_id" validate:"required"`
	Comment string      `json:"comment" form:"comment" validate:"required"`
	Event   EventEntity `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
	User    UserEntity  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
}
