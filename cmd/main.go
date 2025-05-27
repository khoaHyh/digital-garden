package main

import (
	handlers "digital-garden/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	fileServer := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Home)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Default for local development
	}
	port = ":" + port

	logger.Printf("Server starting on port %s", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
