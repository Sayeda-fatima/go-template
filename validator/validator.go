package validator

import "github.com/go-playground/validator"

type Validator interface {
	Validate(input interface{}) error
}
type validateImpl struct {
	validator *validator.Validate
}

// Validate implements Validator.
func (v *validateImpl) Validate(input interface{}) error {
	if err := v.validator.Struct(input); err != nil {
		return err
	}
	return nil
}

func NewValidator() Validator {
	return &validateImpl{
		validator: validator.New(),
	}
}
