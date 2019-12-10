package main

import (
	"github.com/gin-gonic/gin"
	"github.com/material-seminario-golang/deployment/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/material-seminario-golang/deployment/docs"
)

// @title Swagger Example API
// @version 1.0
// @description Sample API to demonstrate how we can generate documentation with Swagger
func main() {
	/*
		TODO: Add Logger
	*/

	router := gin.Default()

	controller := controllers.NewMessageController()

	router.POST("/messages", controller.PostMessage)
	router.GET("/messages", controller.GetMessages)
	router.GET("/messages/:id", controller.GetMessage)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(); err != nil {
		panic(err)
	}
}
