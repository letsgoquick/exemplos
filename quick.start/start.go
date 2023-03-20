package main

import "github.com/jeffotoni/quick"

func main() {
	app := quick.New()
	app.Get("/v1/user", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).SendString("Quick em ação com Cors❤️!")
	})

	app.Get("/v2", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).SendString("Está no ar️!")
	})

	app.Get("/v3", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).SendString("Rodando️!")
	})

	app.Listen("0.0.0.0:8080")
}
