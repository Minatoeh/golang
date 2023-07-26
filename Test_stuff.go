package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// There is every needed components to our blog.
type Blog struct {
	Name      string    `json:"name"`
	Header    string    `json:"header"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var blogRecords []Blog

func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "This is not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Got / request")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Welcome page: Hello, this is my first try to do something in Golang! I hope everything looks fine here.")
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "This is not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Got /blogs request")

	jsonData, err := json.Marshal(blogRecords)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func postBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This is not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Got /post-blog request")

	var newBlog Blog
	err := json.NewDecoder(r.Body).Decode(&newBlog)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	// Perform validation here
	if newBlog.Name == "" || newBlog.Header == "" || newBlog.Content == "" {
		http.Error(w, "Name, Header, and Content fields must not be empty", http.StatusBadRequest)
		return
	}

	newBlog.CreatedAt = time.Now()
	blogRecords = append(blogRecords, newBlog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "You've added a new blog, congrats!"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/blogs", getBlogs)
	mux.HandleFunc("/post-blog", postBlog)

	server := &http.Server{
		Addr:    ":1234",
		Handler: mux,
	}

	fmt.Printf("Server started at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
