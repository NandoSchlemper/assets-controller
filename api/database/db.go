package database;

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)


func loadDb() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	
    if err != nil {
		panic(err)
	}
	
    defer conn.Close(context.Background())
	
    _, err = conn.Exec(context.Background(), 
    "CREATE TABLE IF NOT EXISTS user(id STRING, name STRING)")
	
    if err != nil {
		panic(err)
    }
}
