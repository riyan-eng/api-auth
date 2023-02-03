package util

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Type        string
}

func Validate(input any) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Type = err.Kind().String()
			errors = append(errors, &element)
		}
	}
	return errors
}
