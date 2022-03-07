package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) userHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; cahrset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	user := h.repo.getUser()
	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf("Hello, %s", user.Name),
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error:", err)
	}
}
