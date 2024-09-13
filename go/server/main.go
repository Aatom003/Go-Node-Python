package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Go server on :8080...")
	srv := &http.Server{
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
