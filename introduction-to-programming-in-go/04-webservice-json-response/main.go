package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		result := map[string]string{"name": "Tricia",
			"username": "MacMillan"}
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
