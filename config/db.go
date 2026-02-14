package config

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)

var DB *gorm.DB

func ConnectDB() {
  dsn := "host=localhost user=todo password=pass dbname=todo_db port=5432 sslmode=disable"

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("Failed to connect to database")
  }

  DB = db
}


