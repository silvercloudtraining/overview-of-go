package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		username := r.URL.Query().Get("username")

		w.Header().Add("content-type", "application/json")
		result := map[string]string{"name": name,
			"username": username}
		enc := json.NewEncoder(w)
		err := enc.Encode(result)
		if err != nil {
			log.Print(err)
		}

	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
