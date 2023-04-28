package main

import (
	"log"

	"github.com/jeffotoni/quick"
	"github.com/jeffotoni/quick/middleware/maxbody"
)

func main() {
	app := quick.New()

	app.Use(maxbody.New(0))

	app.Post("/v1/user/maxbody/large", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")

		log.Printf("body: %s", c.BodyString())
		return c.Status(200).Send(c.Body())
	})

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
