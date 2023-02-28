package main

import "github.com/gojeffotoni/quick"

func main() {
	app := quick.New(quick.Config{
		MaxBodySize: 5 * 1024 * 1024,
	})

	/*group := app.Group("/v1")
	group.Get("/user", func(c *quick.Ctx) {
		c.Status(200).SendString("[GET] [GROUP] /v1/user ok!!!")
		return
	})
	group.Post("/user", func(c *quick.Ctx) {
		c.Status(200).SendString("[POST] [GROUP] /v1/user ok!!!")
		return
	})*/

	app.Get("/v2/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com [GET] /v2/user ❤️!")
	})

	app.Post("/v2/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com [POST] /v2/user ❤️!")
	})

	app.Listen("0.0.0.0:8080")
}
