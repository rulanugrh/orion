package middleware

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/rulanugrh/orion/entity/web"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		errors := []web.ValidationList{}
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(data).FieldByName(err.Field())
			errors = append(errors, web.ValidationList{
				Field: field.Type.Name(),
				Error: err.Field() + "|" + err.ActualTag(),
			})
		}

		return web.ValidationError{
			Message: "validation errors",
			Errors:  errors,
		}
	}

	return nil
}
