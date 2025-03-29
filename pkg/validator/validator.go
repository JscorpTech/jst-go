package validator

import (
	"fmt"
	"reflect"

	"github.com/JscorpTech/jst-go/internal/domain/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(st interface{}) []dto.ValidationError {
	validate := validator.New()
	err := validate.Struct(st)
	reflected := reflect.TypeOf(st)
	if reflected.Kind() == reflect.Ptr {
		reflected = reflected.Elem()
	}
	var errors []dto.ValidationError
	var message string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.FieldByName(err.Field())
			fieldName := field.Tag.Get("json")
			if fieldName == "" {
				fieldName = err.Field()
			}
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", fieldName)
			case "email":
				message = fmt.Sprintf("%s is not a valid email address", fieldName)
			default:
				message = fmt.Sprintf("%s is not a valid type", fieldName)
			}
			errors = append(errors, dto.ValidationError{
				Field:   fieldName,
				Type:    err.Tag(),
				Message: message,
			})
		}
		return errors
	}
	return nil
}
