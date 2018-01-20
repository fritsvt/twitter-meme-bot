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
	conn, err := sql.Open("mysql", os.Getenv("DB_USERNAME") + ":"+ os.Getenv("DB_PASSWORD") + "@tcp("+ os.Getenv("DB_HOST") +")/"+os.Getenv("DB_NAME"))
	Connection =  *conn

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
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