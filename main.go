package main

import (
	"github.com/MrRytis/chat-api/internal/router"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	utils.NewDb()
	utils.NewCache()

	router.NewRouter(app)

	app.Listen(":3000")
}
