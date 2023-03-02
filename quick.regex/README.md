# Regex

Regex, ou "Expressões Regulares", é uma técnica utilizada em programação para buscar e manipular padrões de texto. O Framework Quick suporta o uso de regex em rotas, permitindo que desenvolvedores criem rotas dinâmicas e flexíveis.

Para usar regex em rotas no Quick, o desenvolvedor precisa definir uma rota usando uma string que contenha um padrão de expressão regular válido. Isso pode ser feito usando o método HTTP apropriado (Get(), Post(), Put(), etc.) no objeto de aplicativo Quick.

```go
package main

import (
	"github.com/gojeffotoni/quick"
	"github.com/gojeffotoni/quick/middleware/msgid"
)

func main() {
	app := quick.New()

	app.Use(msgid.New())

	app.Get("/v1/user/{[0-9]}", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).String("Quick ação total!!!")
		return
	})

	app.Listen("0.0.0.0:8080")
}
```
