package routes

import (
	"go-echo-template/controller"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, uc controller.UserController) {
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf-token", uc.CsrfToken)
}
