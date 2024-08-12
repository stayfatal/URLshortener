package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager() *DBManager {
	connStr := "user=postgres password=mypass dbname=shortener_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &DBManager{db: db}
}
