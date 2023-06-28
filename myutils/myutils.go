package myutils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres host=localhost password=12345678 dbname=project-db sslmode=disable")
	if err != nil {
		log.Fatal("Cannot start app. Error when connecting to DB: ", err.Error())
	}

	return db
}
