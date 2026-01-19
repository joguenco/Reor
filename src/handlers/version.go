package handlers

import (
	"Reor/src/database"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

type Version struct {
	Name            string `json:"message"`
	Author          string `json:"author"`
	Version         string `json:"version"`
	VersionOS       string `json:"versionOS"`
	VersionLanguage string `json:"versionLanguage"`
	VersionDatabase string `json:"versionDatabase"`
}

func GetVersion(c *fiber.Ctx) error {
	var dbVersion string
	database.ExecuteRawQuery("SELECT version()").Scan(&dbVersion)

	return c.JSON(
		Version{
			Name:            "Reor",
			Author:          "Jorge Luis",
			Version:         "0.0.1",
			VersionOS:       runtime.GOOS + " " + runtime.GOARCH,
			VersionLanguage: runtime.Version(),
			VersionDatabase: dbVersion,
		})
}
