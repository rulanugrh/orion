package domain

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Age      int    `json:"age" form:"age" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Notelp   string `json:"notelp" form:"notelp" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Avatar   string `json:"avatar" form:"avatar" validate:"required"`
}
