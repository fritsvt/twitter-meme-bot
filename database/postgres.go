package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func ConnectPostgres() {
	connStr := "postgres://"+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST"+"/"+os.Getenv("DB_NAME"))

	conn, err := sql.Open("postgres", connStr)
	Connection =  *conn

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
}
x