package mainModule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	userModule "github.com/shinYeongHyeon/eveningIsFoodApi/src/user"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	mainModule := fiber.New()

	mainModule.Mount("/user", userModule.CreateModule())

	// TODO: Error to 404
	mainModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello ✋\nRequest URL %s is not found.", c.Params("*"))
		return c.SendString(msg)
	})

	return mainModule
}
