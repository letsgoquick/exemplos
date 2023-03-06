# GET

O método get pode ser usado para gerar valores aleatórios de diferentes tipos de dados, como inteiros, strings, booleanos, slices, entre outros. Por exemplo, se você estiver testando uma função que recebe um número inteiro como argumento, você pode usar o método get para gerar números aleatórios e verificar se a função está funcionando corretamente para diferentes valores de entrada.

```go
package main

import (
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

	app.Listen("0.0.0.0:8080")
}
```
```go
curl --location 'http://localhost:8080/greet/:name'
--header 'Content-Type", "text/plain' \
--data '
```