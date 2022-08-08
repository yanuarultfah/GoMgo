package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/todo", func(c *fiber.Ctx) error {
		var todo entity.todo
	})
	app.Listen(":3000")
	fmt.Println("Connect GO")
}
