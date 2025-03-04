package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func DbConnect() {
	fmt.Println("Connecting to database...")
	db_string := os.Getenv("DB_STRING_CONNECTION")

	conn, err := pgx.Connect(context.Background(), db_string)
	if err != nil {
		log.Fatal("Unnable to connect to database, error: " + err.Error())
	}
	DB = conn

	fmt.Println("✅ Connection succeded!")

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		status VARCHAR(50) DEFAULT 'active'
	);`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("❌ Error creating tables: %v", err)
	}

	fmt.Println("✅ Tables checked/or created!")
}