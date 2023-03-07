package main

import (
	"fmt"

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

	app.Get("/v2/user", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).SendString("Opa, funcionando!")
	})

	// app.Get("/v3/user", func(c *quick.Ctx) error {
	//	c.Set("Content-Type", "application/json")
	//	return c.Status(200).SendString(c.Query("id"))
	// })

	app.Get("/v1/userx/:p1/:p2/cust/:p3/:p4", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).SendString("Quick em ação ❤️!")
	})

	app.Get("/hello/:name", func(c *quick.Ctx) error {
		name := c.Param("name")
		message := "Olá, " + name + "!"
		return c.Status(200).SendString(message)
	})

	for k, v := range app.GetRoute() {
		fmt.Println(k, "[", v, "]")
	}

	app.Listen("0.0.0.0:8080")
}
