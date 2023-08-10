package main

import (
	"github.com/rulanugrh/orion/handler"
	"github.com/rulanugrh/orion/repository"
	"github.com/rulanugrh/orion/routes"
	"github.com/rulanugrh/orion/services"
)

func main() {
	

	commentRepo := repository.NewCommentRepository()
	commentServ := services.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentServ)

	userRepo := repository.NewUserRepository()
	userServ := services.NewUserServices(userRepo)
	userHandler := handler.NewUserHandler(userServ)

	eventRepo := repository.NewEventRepository()
	eventServ := services.NewEventServices(eventRepo)
	eventHandler := handler.NewEventHandler(eventServ)
	routes.Run(eventHandler, commentHandler, userHandler)
}