package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	rp "real_time_forum/internal/repository"
	"real_time_forum/internal/router"
)

func main() {
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
	fmt.Println("Database setup completed successfully!")

	app_router := router.New_Router(db)

	server := &http.Server{
		Addr:    ":8081",
		Handler: app_router,
	}

	fmt.Println("Server is starting on http://localhost:8081")
	err = server.ListenAndServe();
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
