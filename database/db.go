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
	driver := os.Getenv("DB_DRIVER")

	if driver == "postgres" {
		return PG_GetThreadById(id)
	} else {
		return MYSQL_GetThreadById(id)
	}
}

func InsertThread(thread structs.Thread) {
	driver := os.Getenv("DB_DRIVER")
	if driver == "postgres" {
		PG_InsertThread(thread)
	} else {
		MYSQL_InsertThread(thread)
	}
}