package main

import (
	"log"
	"net/http"
)

func main() {
	filePath := `E:\\Github Projects\\50-golang-python-practices\\50-golang-and-python-practices\\go\\hackernews`
	err := http.ListenAndServe("localhost:8080", http.FileServer(http.Dir(filePath)))
	if err != nil {
		log.Fatalln(err)
	}
}
