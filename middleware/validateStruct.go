package middleware

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		log.Printf("Found Error %v", err.Error())
	}

	return nil
}