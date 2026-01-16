package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the website directory
	fs := http.FileServer(http.Dir("../website"))
	http.Handle("/", fs)

	// Handle button click endpoint
	http.HandleFunc("/api/button-click", handleButtonClick)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleButtonClick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Log to console
	fmt.Println("Button was clicked!")

	// Set CORS headers to allow requests from the website
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	// Send response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Button click received!"}`)
}
