package main

import (
	"github.com/jeffotoni/quick"
)

func main() {
	app := quick.New()

	// Define a rota HTTP GET "/greet/:name"
	app.Get("/greet/:name", func(c *quick.Ctx) error {
		name := c.Param("name")
		c.Set("Content-Type", "text/plain")
		return c.Status(200).SendString("Olá " + name + "!")
	})

	app.Listen("0.0.0.0:8080")
}
