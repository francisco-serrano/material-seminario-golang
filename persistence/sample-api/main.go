package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/material-seminario-golang/persistence/sample-api/controllers"
	"github.com/material-seminario-golang/persistence/sample-api/models"
	"github.com/material-seminario-golang/persistence/sample-api/services"
)

func InitializeDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/sample_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Car{})

	return db
}

func InitializeControllers(db *gorm.DB) controllers.CarController {
	return controllers.CarController{Service: services.NewCarService(db)}
}

func InitializeRoutes(carController controllers.CarController) *gin.Engine {
	router := gin.Default()
	router.GET("/cars/:id", carController.GetCar)
	router.POST("/cars", carController.CreateCar)

	return router
}

func main() {
	db := InitializeDatabaseConnection()
	carController := InitializeControllers(db)
	router := InitializeRoutes(carController)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
