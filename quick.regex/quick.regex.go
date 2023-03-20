package main

import (
	"github.com/jeffotoni/quick"
	"github.com/jeffotoni/quick/middleware/msgid"
)

func main() {
	app := quick.New()

	// adicionando middleware msgid
	app.Use(msgid.New())

	app.Get("/v1/user/{id:[0-9]+}", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).String("Quick ação total!!!")
	})

	app.Use(msgid.New())

	app.Get("/v2/tipos/{id:[0-9]+}", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).String("Quick funcionando!!!")
	})

	app.Listen(":8080")
}
