package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/shinYeongHyeon/eveningIsFood-api/docs"
	mainModule "github.com/shinYeongHyeon/eveningIsFood-api/src"
	"log"
)

// @Title EveningIsFood API
// @Version 0
// @Description 예로부터 약속은 저녁에 이루어 졌고, 역사는 술 속에서 만들어졌습니다.
// @TermsOfService
// @Contact.name EveningIsFood
// @Contact.email eveningisfood@gmail.com
// @Host eveningisfood.com
func main() {
	app := fiber.New(fiber.Config{
		AppName: "Evening is FoodShopCard API 0",
	})

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Mount("/", mainModule.CreateModule())

	log.Fatal(app.Listen(":9999"))
}
