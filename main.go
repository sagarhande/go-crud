package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Variables
type Books struct {
	Id          int
	Title       string
	Author      string
	Description string
}

// Handeler functions
func books(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", books)

	fmt.Println("Server is running!!")
	http.ListenAndServe(":8080", router)
}
