package config

import (
	"api-gin/models"
	"api-gin/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("DATABASE_USERNAME", "root")
	password := utils.Getenv("DATABASE_PASSWORD", "DeNVPix1ZZsXPOFNRp6O")
	host := utils.Getenv("DATABASE_HOST", "containers-us-west-52.railway.app")
	port := utils.Getenv("DATABASE_PORT", "5471")
	database := utils.Getenv("DATABASE_NAME", "railway")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Movie{}, &models.AgeRatingCategory{})

	return db
}
