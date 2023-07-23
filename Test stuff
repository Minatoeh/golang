package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const keyServerAddr = "serverAddr"

type Blog struct {
	Title   string `json:"name"`
	Content string `json:"content"`
	Age     string `json:"age"`
}

var blogRecords []Blog

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))

	io.WriteString(w, "Welcome page: Hello, this is my first try to do something in Golang! I hope you enjoy it, Artem\n")
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
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
