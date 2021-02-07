package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world from fiber!")
	})

	app.Get("/log", func(c *fiber.Ctx) error {
		texto := leerArchivo("server.go")
		return c.SendString(texto)

	})

	app.Listen(":3000")
}

func leerArchivo(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	texto := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		texto += scanner.Text()
		texto += "\n"
	}
	return texto
}
