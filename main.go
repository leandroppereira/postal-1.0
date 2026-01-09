package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type PostalResponse struct {
	PostalID    string `json:"postal_id"`
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
}

func postalHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	resp := PostalResponse{
		PostalID:    id,
		CompanyName: "Royal Mail",
		Phone:       "+44 1234 555555",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] != "--server" {
		log.Println("Use --server to start the service")
		return
	}

	http.HandleFunc("/postal", postalHandler)

	log.Println("Postal backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
