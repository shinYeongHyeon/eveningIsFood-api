package healthControllerQuery

import (
	"github.com/gofiber/fiber/v2"
	"runtime"
)

// HealthCheck : returned 204 HttpStatusCode when server running
func HealthCheck(context *fiber.Ctx) error {
	return context.SendStatus(fiber.StatusNoContent)
}

// GetVersion : returned current Golang version
func GetVersion(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).SendString(runtime.Version())
}
