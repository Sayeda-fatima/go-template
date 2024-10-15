package validator

import (
	"go-echo-template/model"
	"github.com/go-playground/validator"
)

type UserValidator interface {
	UserValidate(user model.User) error
}


type (
	User struct {
		Name string `json:"name" validate:"required" `
		Email string `json:"email" validate:"required, email"`
	}
	userValidator struct {
		validator *validator.Validate
	}
)


func NewUserValidator () UserValidator{
	return &userValidator{
		validator: validator.New(),
	}
}

func (ur *userValidator) UserValidate (user model.User) error {
	if err := ur.validator.Struct(&user); err!=nil{
		return err
	} 
	return nil
}
