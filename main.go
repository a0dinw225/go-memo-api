package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Go API server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go API!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
