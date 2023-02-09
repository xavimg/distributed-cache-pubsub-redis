package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Username string
}

var db *gorm.DB

func Open() error {
	var err error
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		panic("automigrate")
	}

	return nil
}
