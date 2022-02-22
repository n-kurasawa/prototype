package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; cahrset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: "Hello, World.",
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error:", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
