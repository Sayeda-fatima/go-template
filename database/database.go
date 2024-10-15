package database

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	// load env 
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatalf("Error loading .env file")
	  }


	cfg := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	// open connection
	db, err := gorm.Open(mysql.Open(cfg), &gorm.Config{})

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Connected")
	return db;
}
