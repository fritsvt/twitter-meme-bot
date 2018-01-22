package database

import (
	"twitter-meme-bot/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var DB = gorm.DB{}

func Connect() {
	db, err := gorm.Open("sqlite3", os.Getenv("DB_FILE"))
	if err != nil {
		panic("failed to connect database")
	}
	DB = *db;

	// Migrate db
	db.AutoMigrate(&structs.Thread{})
}

func GetThreadById(id string) (bool) {
	var thread structs.Thread
	DB.First(&thread, "reddit_id = ?", id)

	if thread.ID != 0 {
		return true
	}
	return false
}

func InsertThread(thread structs.Thread) {
	DB.Create(&structs.Thread{
		RedditId:thread.RedditId,
		ImageUrl:thread.ImageUrl,
		Title:thread.Title,
		Author:thread.Author,
	})
}