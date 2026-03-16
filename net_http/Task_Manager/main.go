package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID    int    `json:"ID"`
	Title string `json:"Title"`
	Done  bool   `json:"Done"`
}

var (
	tasks []Task
	mu    sync.Mutex
)

func handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		errMsg := fmt.Sprintf("Invalid JSON: %v", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	mu.Lock()
	tasks = append(tasks, t)
	mu.Unlock()
	fmt.Println("Task added to server memory")
	fmt.Fprintf(w, "Task '%s' was successfully added", t.Title)
}

func handleListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Use GET", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Sending tasks to client...")
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed. Use DELETE", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Log: Task %d deleted\n", id)
			fmt.Fprintf(w, "Task %d was successfully deleted", id)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", handleAddTask)
	mux.HandleFunc("/tasks", handleListTasks)
	mux.HandleFunc("/delete", handleDeleteTask)
	fmt.Println("Start server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server died: %v", err)
	}
}
