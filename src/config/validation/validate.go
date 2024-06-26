package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	toolkit "github.com/renatofagalde/golang-toolkit"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateSiteError(validation_err error) *toolkit.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	var restErr toolkit.RestErr

	if errors.As(validation_err, &jsonErr) {
		return restErr.NewBadRequestError("Invalid field type") // json enviando errado
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []restErr.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := restErr.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return restErr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return restErr.NewBadRequestError("Error trying to convert fields")
	}
}
