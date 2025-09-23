package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "postgres.album-solution.svc.cluster.local"
	port = 5432
)

var Db *sql.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = Db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return Db, nil
}
