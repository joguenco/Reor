package main

import (
	"Reor/src/database"
	"Reor/src/handlers"
	"Reor/src/util"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}
	PORT := os.Getenv("PORT")

	app := fiber.New()
	util.ConfigureLogger(app)
	database.ConnectPostgres()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Application": "Reor"})
	})

	app.Get("/ping", handlers.GetPing)
	app.Get("/version", handlers.GetVersion)

	log.Printf("Server is running on port %s\n", PORT)

	app.Listen(":" + PORT)
}
