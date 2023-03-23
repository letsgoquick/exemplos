# Exemplos quick ![Golang Logo](https://raw.githubusercontent.com/jeffotoni/quick/main/quick.png)

### **Bem-vindo ao repositório de exemplos do Quick!**

O **Quick** é um gerenciador de rotas para Go bem flexível e extensível com diversas funcionalidades. O repositório do Framework Quick pode ser encontrado em [aqui](https://github.com/jeffotoni/quick).

Este repositório contém exemplos práticos de como utilizar a biblioteca Quick em Go, uma biblioteca de teste baseada em propriedades que permite escrever testes mais robustos e completos para sua aplicação.

Os exemplos apresentados aqui mostram como utilizar o Quick em diferentes tipos de testes, desde simples até mais complexos, ajudando a começar rapidamente e aprender as melhores práticas de teste.

O Quick é desenvolvido por **jeffotoni** e é uma excelente opção para escrever testes em Go, encontrando falhas em sua aplicação que podem não ser encontradas em testes tradicionais e aumentando a qualidade do código.

Sinta-se à vontade para explorar o repositório, contribuir com seus próprios exemplos e melhorias para a biblioteca Quick. Obrigado por usar Quick!

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

import "github.com/jeffotoni/quick"

func main() {
  app := quick.New()
  app.Get("/v1/user", func(c *quick.Ctx) error {
    c.Set("Content-Type", "application/json")
    return c.Status(200).SendString("Quick em ação com Cors❤️!")
  })
  app.Listen("0.0.0.0:8080")
}

```
