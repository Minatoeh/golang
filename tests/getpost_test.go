// there we will add file and code to our test-check.
package main_test

import (
	"bytes"
	"encoding/json"
	"github.com/Minatoeh/golang/tests"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBlogsEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.getBlogsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, but got %d", rr.Code)
	}
}

func TestPostBlogEndpoint(t *testing.T) {
	title := "Test"
	content := "This this some random stuff for my first golang test. Nothing Special."

	payload := map[string]string{
		"title":   title,
		"content": content,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/blogs", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(main.postBlogHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, but got %d", rr.Code)
	}
}
