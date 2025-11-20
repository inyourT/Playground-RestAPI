package main

import (
	"fmt"
	"playground/internal/handler"
	"playground/internal/repository"
	"playground/internal/server"
	"playground/internal/service"
)

func main() {
	// Init layers
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Init http server
	srv := server.NewServer()

	// Register API routes
	srv.RegisterRouters(userHandler)

	fmt.Println("Server running on Port: 8080")

	srv.Run(":8080")
}