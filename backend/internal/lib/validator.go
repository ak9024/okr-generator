package lib

import (
	"github.com/go-playground/validator"
)

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
