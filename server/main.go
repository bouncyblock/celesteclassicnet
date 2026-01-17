package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/log-message", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// if r.Method == http.MethodOptions {
		// 	w.WriteHeader(http.StatusOK)
		// 	return
		// }

		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }

		var data map[string]string
		json.NewDecoder(r.Body).Decode(&data)

		log.Println("Message from client:", data["message"])

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "logged"})
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
