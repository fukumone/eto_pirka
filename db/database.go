package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func envLoad() {
	if os.Getenv("ETO_PIRKA_ENV") == "" {
		os.Setenv("ETO_PIRKA_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("ETO_PIRKA_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitDB() *gorm.DB {
	envLoad()

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DATABASE_USER_NAME"),
			os.Getenv("DATABASE_USER_PASSWORD"),
			os.Getenv("DATABASE_NAME"),
		))

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
