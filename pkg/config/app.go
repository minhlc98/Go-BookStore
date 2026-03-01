package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	if db != nil {
		return
	}

	dsn := "host=localhost user=admin password=admin dbname=gorm port=5432 sslmode=disable TimeZone=UTC"

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
