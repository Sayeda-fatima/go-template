package main

import (
	"go-echo-template/common"
	"go-echo-template/controller"
	"go-echo-template/database"
	"go-echo-template/middlewares"
	"go-echo-template/repository"
	"go-echo-template/routes"
	"go-echo-template/usecase"
	"go-echo-template/validator"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db := database.NewDB()
	e := echo.New()

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080", os.Getenv("APP_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	common.Newlogger()
	e.Use(middlewares.LoggingMiddleWare)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	userController := controller.NewUserController(userUseCase)
	routes.AuthRoutes(e, userController)
	common.Logger.LogInfo().Msg(e.Start(":8000").Error())
}
