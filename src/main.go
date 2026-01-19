package main

import (
	"Reor/src/common"
	"Reor/src/database"
	"Reor/src/ping"
	receiptModel "Reor/src/receipt/models"
	subscriptionModel "Reor/src/subscription/models"
	"Reor/src/version"
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
	common.ConfigureLogger(app)
	database.ConnectPostgres()
	database.GetDB().Migrator().DropTable(subscriptionModel.Role{}, subscriptionModel.Subscription{}, receiptModel.Taxpayer{}, receiptModel.Receipt{})
	database.GetDB().AutoMigrate(subscriptionModel.Role{})
	database.GetDB().AutoMigrate(subscriptionModel.Subscription{})
	database.GetDB().AutoMigrate(receiptModel.Taxpayer{})
	database.GetDB().AutoMigrate(receiptModel.Receipt{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Application": "Reor"})
	})

	app.Get("/ping", ping.GetPing)
	app.Get("/version", version.GetVersion)

	log.Printf("Server is running on port %s\n", PORT)

	app.Listen(":" + PORT)
}
