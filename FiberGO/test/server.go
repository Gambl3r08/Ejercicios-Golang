package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public/")

	/*app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})*/

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("Value: " + c.Params("value"))
	})

	app.Get("/dato/nombre/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello, " + c.Params("name"))
		}
		return c.SendString("dato vacio")
	})
	app.Listen(":3000")

}
