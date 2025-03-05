package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

func main() {
	db, err := sql.Open("sqlite", "file:data/dev.db?mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = execSqlFile(db, "data/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database initialised")
}

func execSqlFile(db *sql.DB, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return fmt.Errorf("failed to execute SQL: %w", err)
	}

	return nil
}
