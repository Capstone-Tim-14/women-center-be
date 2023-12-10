package helpers

import (
	"reflect"
	"strings"
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
				Field:   strings.ToLower(e.Field()),
				Message: e.Tag(),
			}

			ValidationMessages = append(ValidationMessages, Message)

		}

		return ValidationMessages

	}

	return nil
}

func UniqueDateBooking(fl validator.FieldLevel) bool {

	field := fl.Field()

	if field.Kind() != reflect.Slice {
		return false
	}

	uniqueBookingDate := make(map[string]bool)

	for i := 0; i < field.Len(); i++ {
		val := field.Index(i).Interface().(string)

		if uniqueBookingDate[val] {
			return false
		}
		uniqueBookingDate[val] = true
	}

	return true

}
