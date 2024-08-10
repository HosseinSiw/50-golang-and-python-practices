package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db := NewDb()
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", ShowBooks(db)))
}

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var title, author string
		err := db.QueryRow("SELECT title, author FROM Books").Scan(&title, &author)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(rw, "The first book is '%s' by '%s'", title, author)
	})
}

func NewDb() *sql.DB {
	query := `CREATE TABLE if not exists books(title text, author text,)`
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	return db
}
