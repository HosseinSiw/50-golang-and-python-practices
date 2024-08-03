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
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}

func RetrieveBooks() []Book {
	books := []Book{
		{Title: "Hello", Author: "", NumPages: 1},
		{"Gn", "Girls", 0},
		{"test", "test", 1},
	}
	return books
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	books := RetrieveBooks()

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Cannot find static files")
		return
	}
	err = tmpl.Execute(w, books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
