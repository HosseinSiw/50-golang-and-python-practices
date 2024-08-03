package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type Book struct {
	Title    string
	Author   string
	NumPages int
}

func main() {
	http.HandleFunc("/", ShowBooks)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
func RetrieveBooks() []Book {
	books := []Book{
		{Title: "Hello", Author: "", NumPages: 1},
		{"Gn", "Girls", 0},
		{"test", "test", 1},
		{Title: "GeeksForGeeks", Author: "IDK", NumPages: 4000},
		{Title: "War and Peace", Author: "NULL", NumPages: 800},
	}
	return books
}
func ShowBooks(w http.ResponseWriter, r *http.Request) {
	books := RetrieveBooks()

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
