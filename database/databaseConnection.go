package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github/tiagoduarte/golang-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Inicializa a conexão com o banco de dados
func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresURL := os.Getenv("POSTGRES_URL")

	// Conectando ao banco de dados com GORM
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	// Executa a migração para criar a tabela
	err = db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.Player{},
		&models.Match{},
		&models.MatchScorer{},
		&models.MatchAssistant{},
	)
	if err != nil {
		log.Fatal("Error running migrations:", err)
	}

	// Define a conexão global
	DB = db
}
