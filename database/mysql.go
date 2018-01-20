package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectMysql() {
	conn, err := sql.Open("mysql", os.Getenv("DB_USERNAME") + ":"+ os.Getenv("DB_PASSWORD") + "@tcp("+ os.Getenv("DB_HOST") +")/"+os.Getenv("DB_NAME"))
	Connection =  *conn

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
}
