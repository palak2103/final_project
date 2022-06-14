package main
import (
 	"github.com/go-playground/validator"
)
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// user validation
var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
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

//register validation
var register = validator.New()
func Validate_Struct(registerperson RegisterPerson) []*ErrorResponse {
    var errors []*ErrorResponse
    err := register.Struct(registerperson)
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
