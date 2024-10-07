package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
    validate = validator.New()
	 if validate == nil {
        panic("Validator initialization failed")
    }
}

func ValidateStruct(data interface{}) []string {
	var errors []string
	if validate == nil {
        errors = append(errors, "validator not initialized")
        return errors
    }
    if err := validate.Struct(data); err != nil {
       for _, err := range err.(validator.ValidationErrors) {
            errors = append(errors, fmt.Sprintf("Error in field '%s' with tag '%s'", err.Field(), err.Tag()))
        }
    }
    return errors
}