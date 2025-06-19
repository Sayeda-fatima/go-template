package repository

import (
	"fmt"

	"go-echo-template/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User, jwt string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {

	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}

	return nil

}

func (ur *userRepository) CreateUser(user *model.User) error {

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *model.User, jwt string) error {

	result := ur.db.Model(user).Update("api_token", jwt)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
