package main

import (
	"encoding/json"
	"net/http"
)

type Cat struct {
	Name string `'json:"name"`
	Age int `json:"age"`
	Job string `json: "name"`
}

func catsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
	  cats := []Cat{
		{Name: "Koras02", Age: 27, Job: "Developer"},
		{Name: "PiPi", Age: 22},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cats)
   } else {
       http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
   }
}

func main() {
	http.HandleFunc("/cats", catsHandler)
	http.ListenAndServe(":8080", nil)
}