package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"twitter-meme-bot/structs"
)

var Connection sql.DB

func Connect() {
	driver := os.Getenv("DB_DRIVER")
	if driver == "mysql" {
		ConnectMysql()
	} else if driver == "postgres" {
		ConnectPostgres()
	} else {
		log.Fatal("Invalid database driver")
	}
}

func GetThreadById(id string) (bool) {
	err := Connection.QueryRow("SELECT id FROM tweets WHERE reddit_id=?", id).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Fatal(err)
	default:
		return true // there is a thread
	}
	return false
}

func InsertThread(thread structs.Thread) {
	tx, err := Connection.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO `tweets` (`id`, `reddit_id`, `media_url`, `title`, `author`, `created_at`) VALUES (NULL, (?), (?), (?), (?), CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // danger!
	_, err = stmt.Exec(thread.Id, thread.ImageUrl, thread.Title, thread.Author)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()
}