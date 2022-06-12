package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/shinYeongHyeon/eveningIsFood-api/docs"
	mainModule "github.com/shinYeongHyeon/eveningIsFood-api/src"
	"log"
)

// @title EveningIsFood API
// @version 0
// @description 예로부터 약속은 저녁에 이루어 졌고, 역사는 술 속에서 만들어졌습니다.
// @termsOfService
// @contact.name EveningIsFood
// @contact.email eveningisfood@gmail.com
// @host eveningisfood.com
func main() {
	app := fiber.New(fiber.Config{
		AppName: "Evening is Food API 0",
	})

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Mount("/", mainModule.CreateModule())

	log.Fatal(app.Listen(":9999"))
}
