package main

import (
	// "encoding/json"
	"quran-api-go/router"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// database.ConnectDB()

	var app *fiber.App = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// DisableStartupMessage:        true,
		// DisableDefaultDate:           true,
		// DisableHeaderNormalizing:     true,
		// DisablePreParseMultipartForm: true,
		// DisableDefaultContentType:    true,
		ServerHeader: "biber",
		AppName:      "Quran App v1.0.1",
		ETag:         true,
		GETOnly:      true,
	})

	router.Init(app)

	app.Listen(":3030")
}
