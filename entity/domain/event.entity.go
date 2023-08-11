package domain

import "gorm.io/gorm"

type EventEntity struct {
	gorm.Model
	Name        string              `json:"name" form:"name" validate:"required"`
	Description string              `json:"desc" form:"desc" validate:"required"`
	CreatorID   uint                `json:"user_id" form:"user_id" validate:"required"`
	User        UserEntity          `json:"user" form:"user" gorm:"foreignKey:CreatorID;reference:ID"`
	Participant []ParticipantEntity `json:"participant" form:"participant" gorm:"many2many:parcitipant_event"`
	Comments    []CommentEntity     `json:"comment" form:"comment" gorm:"many2many:comment_event"`
}

type ParticipantEntity struct {
	gorm.Model
	UserID  uint        `json:"user_id" form:"user_id"`
	EventID uint        `json:"event_id" form:"event_id"`
	User    UserEntity  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   EventEntity `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}
