package main

import "github.com/jeffotoni/quick"

func main() {
	app := quick.New(quick.Config{
		MaxBodySize: 5 * 1024 * 1024,
	})

	group := app.Group("/v1")
	group.Get("/user", func(c *quick.Ctx) {
		c.Status(200).SendString("[GET] [GROUP] /v1/user ok!!!")
		return
	})
	group.Post("/user", func(c *quick.Ctx) {
		c.Status(200).SendString("[POST] [GROUP] /v1/user ok!!!")
		return
	})

	group2 := app.Group("/v2")

	group2.Get("/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com [GET] /v2/user ❤️!")
		return
	})

	group2.Post("/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com [POST] /v2/user ❤️!")
		return
	})

	app.Listen("0.0.0.0:8080")
}
