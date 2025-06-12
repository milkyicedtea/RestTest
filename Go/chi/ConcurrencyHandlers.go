package main

import (
	"net/http"
	"time"
)

func HandleConcurrent(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "Bearer test-token" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	time.Sleep(100 * time.Millisecond) // simulate DB or IO

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Success"}`))
}
