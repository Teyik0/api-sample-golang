package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	app.Route("/api", func(api fiber.Router) {
		api.Get("/foo", helloWorld).Name("foo") // /test/foo (name: test.foo)
		api.Get("/bar", bar).Name("bar")        // /test/bar (name: test.bar)
	}, "test.")

	app.Static("/", "./public") // => Serve index.html

	app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) error {
	timestamp := time.Now()
	timestampStr := timestamp.Format("2006-01-02 15:04:05")
	return c.SendString(timestampStr)
}

func bar(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
