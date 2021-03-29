package config

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
	dbport := os.Getenv("dbport")
	database := os.Getenv("database")

	var db *gorm.DB

	if database == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}
		return db
	} else if database == "postgresql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
		db, err := gorm.Open(postgresql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return db

	//connect database

	if err != nil {
		panic("failed to connect database")
	}
	//migrate data models
	db.AutoMigrate(&models.Artikel{})
	return db
}
