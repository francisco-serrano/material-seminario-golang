package mysql_orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Car struct {
	gorm.Model
	Brand      string `gorm:"column:brand;type:varchar(255)"`
	BrandModel string `gorm:"column:model;type:varchar(255)"`
}

func CloseDatabaseConnection(db *gorm.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func Run() {
	db, err := gorm.Open("mysql", "root:root@/sample_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	defer CloseDatabaseConnection(db)

	db.AutoMigrate(&Car{})

	db.Create(&Car{
		Brand:      "vw",
		BrandModel: "golf",
	})

	if err := db.Error; err != nil {
		panic(err)
	}

	var myCar Car
	if err := db.First(&myCar, 1).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", myCar)

	myCar = Car{BrandModel: "golf"}
	if err := db.First(&myCar, "model = ?", "golf").Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", myCar)

	var myCars []Car
	if err := db.Find(&myCars).Error; err != nil {
		panic(err)
	}
	fmt.Println(myCars)
	//fmt.Printf("%+v\n", myCars)

	if err := db.Model(&myCar).Update("BrandModel", "golf gti").Error; err != nil {
		panic(err)
	}

	if err := db.Delete(&myCar).Error; err != nil {
		panic(err)
	}
}
