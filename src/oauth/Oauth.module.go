package oauthModule

import (
	"github.com/gofiber/fiber/v2"
	oauthQueryController "github.com/shinYeongHyeon/eveningIsFood-api/src/oauth/controller/query"
)

func CreateModule() *fiber.App {
	oauthModule := fiber.New()

	oauthModule.Get("/google/signin", oauthQueryController.CreateGoogleSignIn)

	return oauthModule
}
