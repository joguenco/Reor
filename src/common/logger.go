package common

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ConfigureLogger(app *fiber.App) {
	file, err := os.OpenFile("./server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Create a MultiWriter to write to both stdout and the file
	mw := io.MultiWriter(os.Stdout, file)

	// Set the global logger output to the MultiWriter
	log.SetOutput(mw)

	// Configure the internal fasthttp server logger to use the MultiWriter
	// This ensures internal server errors and messages also go to the log file
	app.Server().Logger = log.New(mw, "[Fiber] ", log.LstdFlags)

	app.Use(logger.New(logger.Config{
		Output: mw,
	}))
}
