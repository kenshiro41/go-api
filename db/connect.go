package mydb

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var dbURI = "postgres://ken41:ken41@app_db:5432/mydb?sslmode=disable"

func init() {
	mode := os.Getenv("MODE")
	if mode == "production" {
		dbURI = os.Getenv("DB_URI")
	}

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(dbURI)
	db.LogMode(true)

	DB = db
}
