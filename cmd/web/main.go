package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running!")
	})
	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	fmt.Println("Server is starting on http://localhost:8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}