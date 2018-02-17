package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var connection *sql.DB

func Get() *sql.DB {
	var err error
	connection, err = sql.Open("mysql", "sergey:12345Q@/govideo")
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func Close() {
	if connection != nil {
		connection.Close()
	}
}
