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
	mux.HandleFunc("/toggle-theme", handlers.ToggleTheme)

	// TODO: make env var
	port := ":3000"

	logger.Printf("View your site at http://localhost%s", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Fatalf("Error starting server: %v", err)
	}
}
