package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cecep31/goecho/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Dbcon() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//env
	dbuser := os.Getenv("dbuser")
	dbpass := os.Getenv("dbpass")
	dbname := os.Getenv("dbname")
	dbhost := os.Getenv("dbhost")
	//connect database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//migrate data models
	db.AutoMigrate(&models.Artikel{})
	return db
}
