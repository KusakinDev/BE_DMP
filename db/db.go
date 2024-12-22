package db

import (
	cartstruct "back/struct/cartStruct"
	goodsstruct "back/struct/goodsStruct"
	userstruct "back/struct/userStruct"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=5121508 dbname=dmp_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		log.Fatal(err)
	}
}

func Migration() {
	DB.AutoMigrate(&userstruct.User{})
	DB.AutoMigrate(&goodsstruct.Good{})
	DB.AutoMigrate(&cartstruct.Cart{})
}
