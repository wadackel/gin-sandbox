package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tsuyoshiwada/gin-sandbox/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Get local env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init db
	db, err := gorm.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Tag{})

	// Init router
	router := buildRoutes(db.Debug())
	router.Run(":8080")
}
