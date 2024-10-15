package routes

import (
	"os"

	"go-echo-template/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute(uc controller.UserController) *echo.Echo{
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", os.Getenv("APP_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders,
								echo.HeaderXCSRFToken},	
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},	
		AllowCredentials: true,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	return e
}

