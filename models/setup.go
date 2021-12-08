package models

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "fmt"
)

var DB *gorm.DB

func ConnectDatabase() {


	dbUser := "ofyurgcrjirkqi"
	dbPass := "8a62b1e9bc3ab9d738be35e3d703094c65ffd0531250b22cb4ebffd94b3c69e8"
	dbHost := "ec2-3-95-130-249.compute-1.amazonaws.com"
	dbName := "d1vl98jrcas4kk"
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	
	database.AutoMigrate(&Product{})
	database.AutoMigrate(&Transaction{})
	database.AutoMigrate(&User{})

	DB = database
}
