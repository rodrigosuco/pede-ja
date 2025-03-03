package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func DbConnect() {
	fmt.Println("Connecting to database...")
	db_string := os.Getenv("DB_STRING_CONNECTION")

	_, err := pgx.Connect(context.Background(), db_string)
	if err != nil {
		log.Fatal("Unnable to connect to database, error: " + err.Error())
	} else {
		fmt.Println("Connection succeded!")
	}
}
