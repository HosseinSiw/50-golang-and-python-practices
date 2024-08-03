package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"name"`
	NumPages int    `json:"num_pages"`
}

func main() {
	http.HandleFunc("/books", ShowBooks)
	log.Fatalln(http.ListenAndServe(":7070", nil))
}
func RetrieveBooks() []Book {
	books := []Book{
		{Title: "Hello", Author: "", NumPages: 1},
		{"Gn", "Girls", 0},
		{"test", "test", 1},
	}
	return books
}
func ShowBooks(rw http.ResponseWriter, r *http.Request) {
	books := RetrieveBooks()
	js, err := json.Marshal(&books)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("content/type", "application/json")
	_, err = rw.Write(js)
	if err != nil {
		http.Error(rw, err.Error(), 500)
	}
}
