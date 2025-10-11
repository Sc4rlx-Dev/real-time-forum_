package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	// dbinit "realtimeforum/internal/database/db-init"
	// "realtimeforum/internal/models"
	// "realtimeforum/internal/routes"
	// "realtimeforum/internal/utils"
)
var l = utils.NewRateLimiter()

func init() {
	var err error
	models.Db, err = dbinit.OpenDatabase()
	if err != nil { log.Fatal("Failed to open database:", err)
		return
	}
	err = dbinit.Createtablue(models.Db)
	if err != nil { log.Fatal("Failed to create tables:", err)
		return
	}
	utils.StartSessionCleaner()
}

func main() {
	m := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8081",
		Handler: l.RateMiddleware( routes.Router(m), 20, 100*time.Millisecond, ),
	}
	l.rm_sleep_users()
	server.ListenAndServe()
	fmt.Println("Database setup completed successfully!")
}