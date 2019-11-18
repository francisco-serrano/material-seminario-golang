package models

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	Brand      string `gorm:"column:brand;type:varchar(255)"`
	BrandModel string `gorm:"model:brand;type:varchar(255)"`
	Year       int    `gorm:"year:brand;type:int"`
	EngineType string `gorm:"engine_type:brand;type:varchar(255)"`
}
