package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OPEN_DB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./storage/database.db")
	if err != nil {
		log.Println("Error opening database file:", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	sh, err := os.ReadFile("./migrations/test.sql")
	if err != nil {
		return fmt.Errorf("could not read schema file: %w", err)
	}
	_, err = db.Exec(string(sh))
	if err != nil {
		return fmt.Errorf("error executing schema: %w", err)
	}
	return nil
}
