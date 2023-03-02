# Start

"Start" é um termo geralmente usado em aplicativos web para indicar que o servidor está sendo iniciado e está pronto para lidar com solicitações HTTP. No contexto do Framework Quick, "start" geralmente se refere à ação de iniciar o servidor web Quick.

Para iniciar o servidor web Quick, é preciso criar uma instância do aplicativo Quick, definir as rotas do servidor e chamar a função Listen() passando a porta em que o servidor deve ser executado.

```go
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
```