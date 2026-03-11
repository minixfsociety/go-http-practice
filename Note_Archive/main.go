package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var notes []Note
var nextID int = 1

func listNotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func addNoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newNote Note
	json.NewDecoder(r.Body).Decode(&newNote)
	newNote.ID = nextID
	nextID++
	notes = append(notes, newNote)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newNote)
}

func main() {
	http.HandleFunc("/notes", listNotesHandler)
	http.HandleFunc("/add-note", addNoteHandler)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error:", err)
	}
}
