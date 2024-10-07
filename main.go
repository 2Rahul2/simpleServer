package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{"message": "Hello"}
		json.NewEncoder(w).Encode(response)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
