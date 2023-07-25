// The main goal of this task: I need to make HTTP server with 2 end-points POST and GET.Server should display blog. Each record contains Name, Header, Text and Time
// Get should return json records in blog, Post - you can post new blog. Test it via CURL.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

const keyServerAddr = "serverAddr"

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

	ctx := r.Context()
	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))

	io.WriteString(w, "Welcome page: Hello, this is my first try to do something in Golang! I hope everything looks fine here.\n")
}

// Added Method check for correct request.(If it's not post request , it's should return mistake to client.)
func getBlogs(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "This is not allowed", http.StatusMethodNotAllowed)
		return
	}
	ctx := r.Context()

	fmt.Printf("%s: got /blogs request\n", ctx.Value(keyServerAddr))

	jsonData, err := json.Marshal(blogRecords)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func postBlog(w http.ResponseWriter, r *http.Request) {
	//Check if the method is POST.
	if r.Method != http.MethodPost {
		http.Error(w, "This is not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	fmt.Printf("%s: got /post-blog request\n", ctx.Value(keyServerAddr))

	var newBlog Blog
	err := json.NewDecoder(r.Body).Decode(&newBlog)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %s", err)
		return
	}

	blogRecords = append(blogRecords, newBlog)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "You've added a new blog, congrats!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/blogs", getBlogs)
	mux.HandleFunc("/post-blog", postBlog)

	server := &http.Server{
		Addr:    ":1234",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.WithValue(context.Background(), keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	fmt.Printf("Server started at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
