package domain

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name 		string 	`json:"name" form:"name" validate:"required"`
	Age 		int 	`json:"age" form:"age" validate:"required"`
	Email 		string 	`json:"email" form:"email" validate:"required"`
	Password 	string 	`json:"password" form:"password" validate:"required"`
	Notelp 		string 	`json:"notelp" form:"notelp" validate:"required"`
	Events 		[]ParticipantEntity `json:"event" form:"event" gorm:"many2many:joining_event"`
}