package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := fmt.Fprintln(rw, "Home Page", r.Method)
	if err != nil {
		log.Fatalln(err)
	}
}

func PostIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := fmt.Fprintln(rw, "Post index Page", r.Method)
	if err != nil {
		log.Fatalln(err)
	}
}

func PostCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := fmt.Fprintln(rw, "Post Create page", r.Method)
	if err != nil {
		log.Fatalln(err)
	}
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	_, err := fmt.Fprintln(rw, "Post", id, "page", r.Method)
	if err != nil {
		log.Fatalln(err)
	}
}
func PostUpdateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	_, err := fmt.Fprintln(w, "Updating Post", id, r.Method)
	if err != nil {
		log.Fatalln(err)
	}
}
