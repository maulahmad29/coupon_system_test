package helper

import (
	"coupon_system_test/internal/model/response"

	"github.com/go-playground/validator/v10"
)

func InputValidation(req error) []response.ValidationErrorResponse {
	validationErr := []response.ValidationErrorResponse{}
	for _, errRow := range req.(validator.ValidationErrors) {
		inputAlert := response.ValidationErrorResponse{
			Field: errRow.Field(),
			Tag:   errRow.Tag(),
			Value: errRow.Value(),
			Type:  errRow.Kind().String(),
			Param: errRow.Param(),
		}

		validationErr = append(validationErr, inputAlert)
	}

	return validationErr
}
