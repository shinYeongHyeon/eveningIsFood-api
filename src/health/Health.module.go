package healthModule

import (
	"github.com/gofiber/fiber/v2"
	healthControllerQuery "github.com/shinYeongHyeon/eveningIsFood-api/src/health/controller/query"
)

// CreateModule : returned fiber.App for HealthModule mounting
func CreateModule() *fiber.App {
	healthModule := fiber.New()

	healthModule.Get("/", healthControllerQuery.HealthCheck)
	healthModule.Get("/version", healthControllerQuery.GetVersion)

	return healthModule
}
