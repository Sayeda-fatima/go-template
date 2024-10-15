package main

import (

	"go-echo-template/database"
	"go-echo-template/repository"
	"go-echo-template/usecase"
	"go-echo-template/validator"
	"go-echo-template/routes"
	"go-echo-template/controller"
)

func main(){

	db := database.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUseCase)
	e:= routes.NewRoute(userController)
	e.Logger.Fatal(e.Start(":8000"))
}