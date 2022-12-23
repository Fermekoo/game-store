package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3307)/game_store?parseTime=true"

	DB, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errDb != nil {
		log.Fatalf("DB connection failed: %v", errDb)
	}

	return DB
}
