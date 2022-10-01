package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type ErrorResponse struct {
	Value string `json:"value,omitempty"`
}

func translateError(err error, trans ut.Translator) []string {
	if err == nil {
		return nil
	}
	var errors []string
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		element := e.Translate(trans)
		errors = append(errors, element)
	}
	return errors
}

func Validate(req interface{}) []string {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(req)
	errs := translateError(err, trans)

	return errs
}
