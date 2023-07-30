package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBlogsEndpoint(t *testing.T) {
	req, err := http.NewRequest("POST", "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBlogsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Mistake! Expected status 405 Method Not Allowed, but got %d", rr.Code)
	}

	req, err = http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Mistake! Expected status 404 Method Not Allowed, but got %d", rr.Code)
	}
}

func TestPostBlogEndpoint(t *testing.T) {
	payload := map[string]string{
		"title":   "",
		"content": "",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/blogs", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(postBlogHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Mistake! Expected status 400 Method Not Allowed, but got %d", rr.Code)
	}
}
