package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EchoRequest struct {
	Message string `json:"message"`
}
type EchoResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func echoHendler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request EchoRequest
		json.NewDecoder(r.Body).Decode(&request)
		fmt.Println(request.Message)
		response := EchoResponse{
			Message: request.Message,
			Status:  "success",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
func main() {
	http.HandleFunc("/echo", echoHendler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
