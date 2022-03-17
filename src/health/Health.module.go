package healthModule

import (
	"github.com/gofiber/fiber/v2"
	healthControllerCommand "github.com/shinYeongHyeon/eveningIsFoodApi/src/health/controller/command"
)

// CreateModule : returned fiber.App for HealthModule mounting
func CreateModule() *fiber.App {
	healthModule := fiber.New()

	healthModule.Get("/", healthControllerCommand.HealthCheck)
	healthModule.Get("/version", healthControllerCommand.GetVersion)

	return healthModule
}
