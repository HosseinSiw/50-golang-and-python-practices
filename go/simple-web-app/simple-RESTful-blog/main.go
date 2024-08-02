package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const port = 8080

func main() {
	router := httprouter.New()
	router.GET("/", HomeHandler)

	//	Post signature
	router.GET("/posts", PostIndexHandler)
	router.POST("/posts/create/", PostCreateHandler)

	//	Single post signature
	router.GET("/posts/:id", PostShowHandler)
	router.PUT("/posts/edit/:id/", PostUpdateHandler)

	fmt.Printf("Listen and Serve on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatalln(err)
	}
}
