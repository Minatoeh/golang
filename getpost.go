package main

import (
	"encoding/json"
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

	getBlogsHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Oops! Something went wrong! Expected status 200, but got %d", rr.Code)
	}

	expectedResponse := `[{"title": "Blog ", "content": "Content "}]`
	if rr.Body.String() != expectedResponse {
		t.Errorf("Oops! Something went wrong! Expected response body %s, but got %s", expectedResponse, rr.Body.String())
	}
}
