package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Text string `json:"text"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	response := Message{
		Text: "Hello from Samriddhi AI Backend. This is deployed from Render 🚀",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api", apiHandler)

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
