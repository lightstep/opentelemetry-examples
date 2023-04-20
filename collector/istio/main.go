package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Golang with Istio!")
		log.Println("Request received")
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
