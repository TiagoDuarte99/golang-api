package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBInstance() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PostgresURL := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", PostgresURL)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Testa a conexão com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Could not connect to PostgreSQL:", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	return db
}

// Variável global para a conexão
var DB *sql.DB = DBInstance()

// Função para abrir uma tabela (equivalente a OpenCollection)
func OpenTable(db *sql.DB, tableName string) *sql.DB {
	fmt.Println("Using table:", tableName)
	return db
}