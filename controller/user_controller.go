package controller

import (
	"go-echo-template/common"
	"go-echo-template/dto"
	"go-echo-template/model"
	"go-echo-template/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu                   usecase.UserUsecase
	frequencyRateLimiter *common.FrequencyLimiter
}

func NewUserController(uu usecase.UserUsecase, frequencyRateLimiter *common.FrequencyLimiter) UserController {
	return &userController{uu, frequencyRateLimiter}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid payload data",
			Err:     err.Error(),
		})
	}

	userRes, err := uc.uu.SignUp(user)

	if err != nil {
		return c.JSON(err.ErrorStatusCode(), err)
	}

	return c.JSON(http.StatusCreated, dto.SuccessResponse[model.UserResponse]{
		Message: "Sign up successfull",
		Data:    userRes,
	})
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid payload data",
			Err:     err.Error(),
		})
	}

	if !uc.frequencyRateLimiter.Allow(user.Email) {
		return c.JSON(http.StatusTooManyRequests, map[string]interface{}{
			"message": "You can attempt to login only once every second using this email and password.",
		})

	}

	tokenString, err := uc.uu.Login(user)

	if err != nil {
		return c.JSON(err.ErrorStatusCode(), err)
	}

	return c.JSON(http.StatusOK, dto.SuccessResponse[string]{
		Message: "Login successfull",
		Data:    tokenString,
	})
}

func (uc *userController) Logout(c echo.Context) error {

	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid payload data",
			Err:     err.Error(),
		})
	}

	if err := uc.uu.Logout(user); err != nil {
		return c.JSON(err.ErrorStatusCode(), err)
	}

	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {

	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
