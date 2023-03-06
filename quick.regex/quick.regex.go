package main

import (
	"github.com/jeffotoni/quick"
	"github.com/jeffotoni/quick/middleware/msgid"
)

func main() {
	app := quick.New()

	app.Use(msgid.New())

	app.Get("/v1/user/{id:[0-9]+}", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Status(200).String("Quick ação total!!!")
	})

	app.Listen("0.0.0.0:8080")
}
