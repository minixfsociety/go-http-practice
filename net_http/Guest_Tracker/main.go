package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Enter struct {
	Name string `json:"name"`
}

func handleEnter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Choose the correct method!", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "I am waiting for a JSON", http.StatusBadRequest)
		return
	}
	var e Enter
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Error in the JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Guest %v registered!", e.Name)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/enter", handleEnter)
	fmt.Println("Start server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error:", err)
	}
}
