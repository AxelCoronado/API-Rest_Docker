package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN="host=go_db user=axel password=apipostgres dbname=gorm port=5432"
var DB *gorm.DB

func DBconnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}