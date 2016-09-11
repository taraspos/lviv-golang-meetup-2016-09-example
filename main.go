package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, str)
		log.Println(str)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Golang Meetup!")
		log.Println("Hello, Golang Meetup")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
