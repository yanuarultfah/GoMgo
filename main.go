package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/:name", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello " + ctx.Params("name"))
	})
	app.Listen(":3000")

}
