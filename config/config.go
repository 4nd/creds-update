package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func GetDatabase() *sql.DB {

	dbConfig := mysql.Config{
		User:      "root",
		Passwd:    "",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "credentials",
		ParseTime: true,
	}

	var db, err = sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
