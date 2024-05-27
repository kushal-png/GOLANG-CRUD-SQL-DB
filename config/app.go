package config

import (
	"fmt"
	"goserver/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
const DNS = "root:vivov3designed@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect withb Db")
		log.Fatal(err)
		return
	}

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Failed to migrate database")
		log.Fatal(err)
	}

	fmt.Println("Database connection and migration successful")
	
}

func GetDb() *gorm.DB {
	return DB
}
