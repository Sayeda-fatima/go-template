package controller

import (
	"go-echo-template/common"
	"go-echo-template/model"
	"go-echo-template/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu          usecase.UserUsecase
	rateLimiter *common.RateLimiter
}

func NewUserController(uu usecase.UserUsecase, rateLimiter *common.RateLimiter) UserController {
	return &userController{uu, rateLimiter}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.SignUp(user)

	if err != nil {
		return c.JSON(err.ErrorStatusCode(), err)
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if !uc.rateLimiter.Allow(user.Email, 1*time.Second) {
		return c.JSON(http.StatusTooManyRequests, map[string]interface{}{
			"message": "You can attempt to login only once every second using this email and password.",
		})

	}

	tokenString, err := uc.uu.Login(user)

	if err != nil {
		return c.JSON(err.ErrorStatusCode(), err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (uc *userController) Logout(c echo.Context) error {

	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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
