package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    rp "real_time_forum/internal/repository"
)

func main() {
    //  <---DATABASE INIT---->
    if err := os.MkdirAll("storage", os.ModePerm); err != nil {
        log.Fatalf("Failed to create storage directory: %v", err)
    }

    db, err := rp.OPEN_DB()
    if err != nil {
        log.Fatal("Failed to open database:", err)
    }
	
	if err := rp.CreateTables(db); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Server is running!")
    })

    server := &http.Server{
        Addr:    ":8081",
        Handler: mux,
    }

    fmt.Println("Server is starting on http://localhost:8081")
    err = server.ListenAndServe()
    if err != nil {
        log.Fatal("Server failed to start:", err)
    }
}