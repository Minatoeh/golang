package main

import (
	"encoding"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getBlogsHandler(w http.ResponseWriter, r *http.Request) {
	titles := []map[string]string{
		{"title": "name ", "content": "Content"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(titles)
}
