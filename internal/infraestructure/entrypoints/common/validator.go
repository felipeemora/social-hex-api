package entrypoints

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

var ValidationMessages = map[string]string{
	"required": "This field is required",
	"min":      "This field must have at least %s characters",
	"max":      "This field must have at most %s characters",
}

func ValidateStruct(s any) map[string]any {
	errorMessages := make(map[string]any)

	if err := Validate.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			messageTemplate, exists := ValidationMessages[fieldError.Tag()]
			if !exists {
				messageTemplate = "Invalid value"
			}
			errorMessages[fieldError.Field()] = fmt.Sprintf(messageTemplate, fieldError.Param())
		}
	}
	return errorMessages
}
