package config

import (
	"fmt"
	"log"
	"os"

	"github.com/cecep31/goecho/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
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
<<<<<<< HEAD

=======
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",dbhost, dbuser, dbpass, dbname, dbport)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
>>>>>>> 71dd8b9488b9963988b9f6ecc2153ec516dfc548
	if err != nil {
		panic("failed to connect database")
	}
	//migrate data models
	db.AutoMigrate(&models.Artikel{})
	return db
}
