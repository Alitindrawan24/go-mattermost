package server

import (
	"fmt"

	"github.com/Alitindrawan24/go-mattermost/internal/app"
	"github.com/gin-gonic/gin"
)

func Run() {
	repository := app.NewRepository()
	service := app.NewService(repository)
	handler := app.NewHandler(service)

	router := gin.Default()
	router.GET("/", handler.User.HandleGetUser)
	router.POST("/custom-status", handler.CustomStatus.HandleSetCustomStatus)

	port := 8080
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server: %v", err))
	}
}
