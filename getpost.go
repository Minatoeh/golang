package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// added t.Run function
func TestHandler(t *testing.T) {
	t.Run("GetBlogsEndpoint", testGetBlogsEndpoint)
	t.Run("PostBlogEndpoint", testPostBlogEndpoint)
}

func testGetBlogsEndpoint(t *testing.T) {
	t.Run("ValidGetRequest", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/blogs", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getBlogsHandler)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200 OK, but got %d", rr.Code)
		}
	})

	t.Run("NonExistentEndpoint", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/nonexistent", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(getBlogsHandler)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected status 404 Not Found, but got %d", rr.Code)
		}
	})
}

func testPostBlogEndpoint(t *testing.T) {
	t.Run("ValidPostRequest", func(t *testing.T) {
		payload := map[string]string{
			"title":   "Test",
			"content": "This this some random stuff for my first golang test.Nothing Special.",
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

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status 201 Created, but got %d", rr.Code)
		}
	})

	t.Run("InvalidPayload", func(t *testing.T) {
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
			t.Errorf("Expected status 400 Bad Request, but got %d", rr.Code)
		}
	})
}
