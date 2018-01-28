package database

import (
	"twitter-meme-bot/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"log"
)

var DB = gorm.DB{}

func Connect() {
	driver := os.Getenv("DB_CONNECTION")

	if driver == "mysql" {
		db, err := gorm.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect database")
		}
		DB = *db;
	} else if driver == "postgres" {
		db, err := gorm.Open("postgres", "host="+os.Getenv("DB_HOST")+" user="+os.Getenv("DB_USERNAME")+" dbname="+os.Getenv("DB_DATABASE")+" sslmode=disable password="+os.Getenv("DB_PASSWORD"))
		if err != nil {
			panic("failed to connect database")
		}
		DB = *db;
	} else if driver == "sqlite" {
		db, err := gorm.Open("sqlite3", os.Getenv("DB_FILE"))
		if err != nil {
			panic("failed to connect database")
		}
		DB = *db;
	} else {
		log.Fatal("Database driver not found")
	}

	// Migrate db
	DB.AutoMigrate(&structs.Thread{}, &structs.ScheduledTweet{})
}