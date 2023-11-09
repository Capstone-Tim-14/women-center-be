package helpers

import (
	"woman-center-be/utils/exceptions"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidationError(ctx echo.Context, err error) []exceptions.ValidationMessage {
	validationError, ok := err.(validator.ValidationErrors)
	if ok {

		ValidationMessages := []exceptions.ValidationMessage{}

		for _, e := range validationError {

			Message := exceptions.ValidationMessage{
				Field:   e.Field(),
				Message: e.Tag(),
			}

			ValidationMessages = append(ValidationMessages, Message)

		}

		return ValidationMessages

	}

	return nil
}
