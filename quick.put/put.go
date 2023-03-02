package main

import (
	"github.com/gojeffotoni/quick"
)

func main() {
	app := quick.New()

	app.Put("/users/:id", func(c *quick.Ctx) {
		userID := c.Param("id")
		// Lógica de atualização do usuário
		c.SendString("Usuário " + userID + " atualizado com sucesso!")
	})

	app.Listen(":8080")
}
