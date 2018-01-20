package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"twitter-meme-bot/structs"
)

func ConnectMysql() {
	conn, err := sql.Open("mysql", os.Getenv("DB_USERNAME") + ":"+ os.Getenv("DB_PASSWORD") + "@tcp("+ os.Getenv("DB_HOST") +")/"+os.Getenv("DB_NAME"))
	Connection =  *conn

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
}

func MYSQL_GetThreadById(id string) (bool) {
	query := "SELECT id FROM tweets WHERE reddit_id= ?"
	err := Connection.QueryRow(query, id).Scan(&id)
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


func MYSQL_InsertThread(thread structs.Thread) {
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