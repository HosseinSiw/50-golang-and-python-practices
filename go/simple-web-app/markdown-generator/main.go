package main

import (
	"github.com/russross/blackfriday"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe("localhost:6060", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownBasic([]byte(r.FormValue("body")))
	_, err := rw.Write(markdown)
	if err != nil {
		log.Fatalln(err)
	}
}
