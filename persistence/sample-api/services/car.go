package services

import (
	"github.com/jinzhu/gorm"
	"github.com/material-seminario-golang/persistence/sample-api/models"
	"github.com/material-seminario-golang/persistence/sample-api/views"
	"strconv"
)

type CarService struct {
	db *gorm.DB
}

func NewCarService(db *gorm.DB) *CarService {
	return &CarService{
		db: db,
	}
}

func (c *CarService) GetCar(idStr string) (*models.Car, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	var car models.Car

	if err := c.db.Find(&car, id).Error; err != nil {
		return nil, err
	}

	return &car, nil
}

func (c *CarService) CreateCar(car views.CreateCarRequest) (*models.Car, error) {
	myCar := models.Car{
		Brand:      car.Brand,
		BrandModel: car.Model,
		Year:       car.Year,
		EngineType: car.EngineType,
	}

	if err := c.db.Create(&myCar).Error; err != nil {
		return nil, err
	}

	return &myCar, nil
}
