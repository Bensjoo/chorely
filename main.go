package main

import (
	"fmt"
	"log"
	"net/http"

	"chorely/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/hello", handlers.Hello)

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
