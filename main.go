package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:name", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello GO")
	})
	app.Listen(":3000")
	fmt.Println("Connect GO")
}
