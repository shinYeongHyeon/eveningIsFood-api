package main

import (
	"github.com/gofiber/fiber/v2"
	mainModule "github.com/shinYeongHyeon/eveningIsFood-api/src"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Evening is Food API v1.0.0-pre",
	})

	app.Mount("/", mainModule.CreateModule())

	log.Fatal(app.Listen(":9999"))
}
