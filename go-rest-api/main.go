package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greetings", GreetingsHandler).Methods("GET")
	http.Handle("/", r)
	println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func GreetingsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is missing", http.StatusBadRequest)
		return
	}
	response := map[string]string{"message": "Hello " + name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}