package utils

import (
	"fmt"
	"log"

	"github.com/amanuel15/fiber_server/pkg/interfaces"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBody(data interface{}) interface{} {
	log.Println("Vlidate; ", data)

	err := validate.Struct(data)
	if err != nil {
		var response interfaces.IResponse
		for _, err := range err.(validator.ValidationErrors) {
			response.Error = fmt.Sprintf("\"%s\" is %s but found \"%s\"", err.Field(), err.Tag(), err.Param())
			return response
		}
	}
	return nil
}
