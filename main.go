package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Radupultea1"
	dbname   = "practica"
)

type Catalog struct {
	db *sql.DB
}

func main() {
	pgsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pgsqlInfo)

	if err != nil {
		panic(err)
	}
	catalog := &Catalog{db: db}

	http.HandleFunc("/", serveLogin)
	http.HandleFunc("/login", catalog.loginHandler)

	http.ListenAndServe(":8080", nil)
}
