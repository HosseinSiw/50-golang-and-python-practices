package main

import (
	"log"
	//"fmt"

	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {

	n := negroni.New(
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewRecovery(),
		negroni.NewStatic(http.Dir("public")),
	)
	log.Fatalln(http.ListenAndServe(":6060", n))
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	str := r.URL.Query().Get("password")
	log.Println(r.URL.Query())
	_, err := rw.Write([]byte(str))
	if err != nil {
		log.Fatalln(err)
	}
}
