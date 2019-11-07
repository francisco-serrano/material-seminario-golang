package main

import (
	"github.com/gin-gonic/gin"
	"github.com/material-seminario-golang/rest/sampleapi/controllers"
)


func main() {
	router := gin.Default()

	controller := controllers.NewMessageController()

	router.POST("/messages", controller.PostMessage)
	router.GET("/messages", controller.GetMessages)
	router.GET("/messages/:id", controller.GetMessage)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
