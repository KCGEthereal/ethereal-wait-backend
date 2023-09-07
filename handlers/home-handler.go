package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(map[string]string{
		"hello": "world",
	}); err != nil {
		log.Println("Unable to print to browser")
	}
}
