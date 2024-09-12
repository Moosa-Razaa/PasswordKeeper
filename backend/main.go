package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s %s", r.RemoteAddr, r.URL.Path)
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
