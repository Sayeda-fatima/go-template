package usecase

import (
	"net/http"
	"os"
	"time"

	"go-echo-template/common"
	"go-echo-template/constants"
	"go-echo-template/model"
	"go-echo-template/repository"
	"go-echo-template/validator"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecase interface {
		SignUp(user model.User) (model.UserResponse, common.HttpError)
		Login(user model.User) (string, common.HttpError)
		Logout(user model.User) common.HttpError
	}

	userUsecase struct {
		ur repository.UserRepository
		uv validator.UserValidator
	}
	JwtCustomClaims struct {
		id string `json:"id"`
		jwt.RegisteredClaims
	}
)

func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, common.HttpError) {

	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, common.NewHTTPError(http.StatusUnprocessableEntity, constants.MsgUnprocessableEntity, err.Error())

	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, common.NewHTTPError(http.StatusInternalServerError, constants.MsgInternalServerError, err.Error())

	}

	newUser := model.User{Email: user.Email, Name: user.Name, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, common.NewHTTPError(http.StatusInternalServerError, constants.MsgInternalServerError, err.Error())

	}

	resUser := model.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, common.HttpError) {

	if err := uu.uv.UserValidate(user); err != nil {
		common.Logger.LogError().Msg(err.Error())

		return "", common.NewHTTPError(http.StatusUnprocessableEntity, constants.MsgUnprocessableEntity, err.Error())

	}

	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", common.NewHTTPError(http.StatusBadRequest, constants.MsgInvalidCredentials, err.Error())

	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", common.NewHTTPError(http.StatusInternalServerError, constants.MsgInternalServerError, err.Error())

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  storedUser.ID,
		"exp": time.Now().Add(time.Hour * 1000000).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", common.NewHTTPError(http.StatusInternalServerError, constants.MsgInternalServerError, err.Error())

	}
	// store jwt token to db
	if err := uu.ur.UpdateUser(&storedUser, tokenString); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return "", common.NewHTTPError(http.StatusInternalServerError, constants.MsgInternalServerError, err.Error())

	}
	return tokenString, nil
}

func (uu *userUsecase) Logout(user model.User) common.HttpError {

	storedUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return common.NewHTTPError(http.StatusBadRequest, constants.MsgInvalidEmail, err.Error())
	}
	if err := uu.ur.UpdateUser(&storedUser, ""); err != nil {
		common.Logger.LogError().Msg(err.Error())
		return common.NewHTTPError(http.StatusInternalServerError, constants.MsgInternalServerError, err.Error())
	}

	return nil
}
