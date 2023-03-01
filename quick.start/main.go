package main

import "github.com/gojeffotoni/quick"

func QuickHandler(c *quick.Ctx) {
	c.Set("Content-Type", "application/json")

	type my struct {
		Msg string `json:"msg"`
		Key string `json:"key"`
		Val string `json:"val"`
	}

	c.Status(200).JSON(&my{
		Msg: "Quick ❤️",
		Key: c.Param("param1"),
		Val: c.Param("param2"),
	})
}

func main() {
	app := quick.New()

	app.Get("/v1/customer", QuickHandler)

	app.Listen("0.0.0.0:8080")
}
