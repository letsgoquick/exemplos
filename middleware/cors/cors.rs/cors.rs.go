package main

import (
	"fmt"
	"log"

	"github.com/gojeffotoni/quick"
	"github.com/gojeffotoni/quick/middleware/msgid"
	"github.com/rs/cors"
)

func main() {

	app := quick.New()

	app.Post("/v1/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		type My struct {
			Name string `json:"name"`
			Year int    `json:"year"`
		}

		var my My
		err := c.BodyParser(&my)
		fmt.Println("byte:", c.Body())

		if err != nil {
			c.Status(400).SendString(err.Error())
			return
		}

		fmt.Println("String:", c.BodyString())
		c.Status(200).JSON(&my)
		return
	})

	app.Use(msgid.New())
	app.Get("/v1/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).String("Quick ação total!!!")
		return
	})

	mux := cors.Default().Handler(app)
	log.Fatal(app.Listen("0.0.0.0:8080", mux))
}
