package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	_db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=postgres dbname=gin port=5432 sslmode=disable TimeZone=Asia/Almaty",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := _db.AutoMigrate(&Grocery{}); err != nil {
		log.Fatal(err.Error())
	}

	DB = _db
}
