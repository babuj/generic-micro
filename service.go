package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "6060"
	}
	return ":" + port
}

func main() {
	fmt.Println("Starting the hello Greet service at " + port() + " port.")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(port(), router))
}
