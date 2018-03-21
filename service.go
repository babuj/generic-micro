package main

import (
	"fmt"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Name of the service : "+name()+"\n")
}

func Echo(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query()["text"][0]
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text-plain")
	fmt.Fprintf(w, text)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "6060"
	}
	return ":" + port
}

func name() string {
	name := os.Getenv("NAME")
	if len(name) == 0 {
		name = "echo"
	}
	return name
}

func main() {
	fmt.Println("Starting the " + name() + " service at " + port() + " port.")
	http.HandleFunc("/", Index)
	http.HandleFunc("/echo", Echo)
	http.ListenAndServe(port(), nil)
}
