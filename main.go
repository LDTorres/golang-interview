package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/LDTorres/golang-interview/internal/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get DB credentials
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping database
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	app := fiber.New()

	http.SetupRoutes(app, db)

	app.Listen(":3000")
}
