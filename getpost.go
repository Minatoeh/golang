package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBlogsEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBlogsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Oops! Something went wrong! Expected status 200, but got %d", rr.Code)
	}

	expectedResponse := `[{"title": "Blog ", "content": "Content "}]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Oops! Something went wrong! Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}

func TestPostBlogEndpoint(t *testing.T) {
	payload := map[string]string{
		"title":   "Test Blog",
		"content": "Test Content",
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

	if rr.Code != http.StatusOK {
		t.Errorf("Oops! Something went wrong! Expected status 200, but got %d", rr.Code)
	}
}
