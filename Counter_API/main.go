package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var count int

type AddRequest struct {
	Amount int `json:"Amount"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Нужен GET запрос", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Текущее значение счетчика: %d", count)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен POST запрос", http.StatusMethodNotAllowed)
		return
	}
	var request AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Ошибка в JSON", http.StatusBadRequest)
		return
	}
	count += request.Amount
	fmt.Fprintf(w, "Добавлено: %d. Теперь счетчик: %d", request.Amount, count)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Нужен POST запрос", http.StatusMethodNotAllowed)
		return
	}
	count = 0
	fmt.Fprint(w, "Счетчик сброшен")
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/reset", resetHandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
