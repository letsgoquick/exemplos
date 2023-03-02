# Examples

Bem-vindo ao repositório de exemplos do Quick!

Este repositório contém exemplos práticos e úteis de como utilizar a biblioteca Quick em Go. Quick é uma biblioteca de teste de propriedade baseada em propriedades que permite escrever testes mais robustos e completos para sua aplicação.

Aqui você encontrará exemplos de como utilizar Quick em diferentes tipos de testes, desde testes simples até testes mais complexos. Esses exemplos são projetados para ajudá-lo a começar rapidamente com Quick e aprender as melhores práticas de teste.

A biblioteca Quick é desenvolvida por **gojeffotoni** e é uma excelente opção para escrever testes em Go. Ela ajuda a encontrar falhas em sua aplicação que podem não ser encontradas em testes tradicionais, aumentando assim a qualidade do código.

Esperamos que estes exemplos sejam úteis para você e ajudem a melhorar a qualidade de seus testes. Sinta-se à vontade para explorar o repositório e contribuir com seus próprios exemplos e melhorias para a biblioteca Quick. Obrigado por usar Quick!

## Quais exemplos você encontrará no repositório?

  * [Group](/group/)
  * [Middleware](/middleware/)
  * [Quick.delete](quick.delete/)
  * [Quick.post](quick.post/)
  * [Quick.put](quick.put/)
  * [Quick.regex](quick.regex/)
  * [Quick.start](quick.start/)

```go
package main

import "github.com/gojeffotoni/quick"
import "github.com/gojeffotoni/quick/middleware/cors"

func main() {
  app := quick.New()
  app.Get("/v1/user", func(c *quick.Ctx) {
    c.Set("Content-Type", "application/json")
    c.Status(200).SendString("Quick em ação com Cors❤️!")
  })
  app.Listen("0.0.0.0:8080")
}

```
