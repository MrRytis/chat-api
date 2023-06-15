package main

import (
	"database/sql"
	"fmt"
	"github.com/MrRytis/chat-api/internal/migrations"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func main() {
	color.Green("Running migration down")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	var db *sql.DB
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open: ", err)
	}
	defer db.Close()

	goose.SetBaseFS(migrations.EmbedMigrationsFs)

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal("Failed to set dialect: ", err)
	}

	if err = goose.Up(db, "migrations"); err != nil {
		log.Fatal("Failed migration: ", err)
	}

	color.Green("Migration ran successfully")
}
