package main

import (
	_ "github.com/MrRytis/chat-api/docs"
	"github.com/MrRytis/chat-api/internal/router"
	"github.com/MrRytis/chat-api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
)

// @title Chat applications API
// @version 1.0
// @description This API is used for chat application
// @contact.name Rytis
// @contact.email rytis.janceris@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	utils.NewDb()
	utils.NewCache()

	app.Get("/*", swagger.HandlerDefault)

	router.NewRouter(app)

	app.Listen(":3000")
}
