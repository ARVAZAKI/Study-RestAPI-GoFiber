package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	fmt.Println("Server is running on port 8080")
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}

}
