package config

import (
  "fmt"
  "log"
  "os"

  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "github.com/joho/godotenv"
)

func env_load() {
  if os.Getenv("ETO_PIRKA_ENV") == "" {
    os.Setenv("ETO_PIRKA_ENV", "development")
  }

  err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("ETO_PIRKA_ENV")))
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func Database() gorm.DB {
  env_load()
  db, err := gorm.Open("mysql",
    fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True&loc=Local",
      os.Getenv("DB_USER_NAME"),
      os.Getenv("DATABASE_NAME"),
    ))

  if err != nil {
    panic("failed to connect database")
  }
  return db
}
