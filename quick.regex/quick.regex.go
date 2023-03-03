package main

import (
	"github.com/gojeffotoni/quick"
	"github.com/gojeffotoni/quick/middleware/msgid"
)

func main() {
	app := quick.New()

	app.Use(msgid.New())

	app.Get("/v1/user/{id:[0-9]+}", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).String("Quick ação total!!!")
		return
	})

	app.Listen("0.0.0.0:8080")
}
