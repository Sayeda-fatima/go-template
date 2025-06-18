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
	"net/http"
	"os"

	"go-echo-template/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	envErr := godotenv.Load(".env")
	if envErr != nil {
		common.Logger.LogError().Msg("Error loading .env file")
		return
	}
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

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
	common.NewLogger(os.Getenv("ENV"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*", "http://localhost:8080", os.Getenv("APP_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, "Authorization"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowCredentials: false,
	}))

	e.Use(middlewares.LoggingMiddleWare)
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         4 * 1024,
		LogLevel:          log.ERROR,
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
				"error":   err.Error(),
				"stack":   string(stack),
			})
		},
	}))
	userController := controller.NewUserController(userUseCase, common.NewRateLimiter())
	routes.AuthRoutes(e, userController)
	common.Logger.LogInfo().Msg(e.Start(":8000").Error())
}
