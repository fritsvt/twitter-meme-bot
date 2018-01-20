package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"log"
	"twitter-meme-bot/structs"
)

func ConnectPostgres() {
	connStr := "postgres://"+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST"+"/"+os.Getenv("DB_NAME"))
	if os.Getenv("DATABASE_URL") != "" {
		connStr = os.Getenv("DATABASE_URL")
	}
	conn, err := sql.Open("postgres", connStr)
	Connection =  *conn

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
}

func PG_GetThreadById(id string) (bool) {
	query:= "SELECT id FROM tweets WHERE reddit_id= '"+id+"'"

	err := Connection.QueryRow(query).Scan(&id)
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

func PG_InsertThread(thread structs.Thread) {
	sqlStatement := "INSERT INTO public.tweets(reddit_id, media_url, title, author) VALUES($1, $2, $3, $4) RETURNING 'id', 'reddit_id', 'media_url', 'title', 'author', 'created_at'"
	_, err := Connection.Exec(sqlStatement, thread.Id, thread.ImageUrl, thread.Title, thread.Author)
	if err != nil {
		panic(err)
	}
}