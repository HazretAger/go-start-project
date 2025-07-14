package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitSchema(db *sql.DB) {
    stmts := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            surname TEXT NOT NULL,
            middle_name TEXT NOT NULL,
            birth_date DATETIME NOT NULL,
            phone_number TEXT NOT NULL,
            email TEXT NOT NULL,
            password TEXT NOT NULL,
            confirm_password TEXT NOT NULL,
            is_verified BOOLEAN NOT NULL DEFAULT FALSE,
            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        );`,
        // добавляй сюда другие таблицы
    }

    for _, stmt := range stmts {
        if _, err := db.Exec(stmt); err != nil {
            log.Fatalf("Failed to create schema: %v\nSQL: %s", err, stmt)
        }
    }

    log.Println("Database schema initialized")
}

func Connect(dbPath string) *sql.DB {
    db, err := sql.Open("sqlite", dbPath)
	
    if err != nil {
        log.Fatal("Cannot open db:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Cannot ping db:", err)
    }

    return db
}