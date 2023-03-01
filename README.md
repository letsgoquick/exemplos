# Examples

### Quick

Quick é um framework web para a linguagem Go que permite criar aplicativos web de forma rápida e fácil. Com uma sintaxe simples e concisa, Quick torna fácil lidar com solicitações HTTP e criar respostas dinâmicas.

O repositório [https://github.com/gojeffotoni/quick](https://github.com/gojeffotoni/quick) é um ótimo exemplo de como usar o Quick para criar aplicativos web. A seguir, separamos alguns exemplos em categorias para ilustrar as possibilidades oferecidas pela biblioteca.

### Rotas

A biblioteca [https://github.com/gojeffotoni/quick](https://github.com/gojeffotoni/quick) inclui vários exemplos de como definir rotas em seu aplicativo Quick. O exemplo abaixo define uma rota para lidar com solicitações GET para `/v1/customer`:

```go
app.Get("/v1/customer", QuickHandler)
```

Neste exemplo, `app` é uma instância do objeto `quick.App`, e `QuickHandler` é uma função que será executada sempre que uma solicitação GET para `/v1/customer` for recebida.

### Manipuladores de solicitação

O exemplo [https://github.com/gojeffotoni/quick](https://github.com/gojeffotoni/quick) inclui vários exemplos de como definir manipuladores de solicitação em seu aplicativo Quick. O exemplo abaixo define um manipulador de solicitação que retorna uma resposta JSON com uma mensagem personalizada, bem como dois parâmetros de rota `param1` e `param2`:

```go
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
```

Neste exemplo, c é uma instância do objeto `quick.Ctx`, que fornece acesso a informações sobre a solicitação HTTP recebida. O manipulador define um tipo my que será serializado como uma resposta JSON, e usa o método c.Param para obter os valores dos parâmetros de rota `param1` e `param2`.
Respostas

A biblioteca [https://github.com/gojeffotoni/quick](https://github.com/gojeffotoni/quick) inclui vários exemplos de como criar respostas em seu aplicativo Quick. O exemplo abaixo cria uma resposta HTML simples:

```go

func QuickHandler(c *quick.Ctx) {
    c.Status(200).HTML("<h1>Hello, world!</h1>")
}

```

Neste exemplo, o manipulador de solicitação chama o método `c.Status` para definir o código de status HTTP da resposta, e então chama o método `c.HTML` para definir o conteúdo HTML da resposta.

### Middleware

A biblioteca [https://github.com/gojeffotoni/quick](https://github.com/gojeffotoni/quick) inclui vários exemplos de como usar middleware em seu aplicativo Quick. O exemplo abaixo usa o middleware Logger para registrar informações sobre cada solicitação recebida:

```go

func main() {
    app := quick.New()

    app.Use(quick.Logger())

    // ...rotas e manipuladores de solicitação aqui...

    app.Listen("0.0.0.0:8080")
}
```

Neste exemplo, app.Use é usado para adicionar o middleware Logger à cadeia de middleware do aplicativo Quick. Isso significa que o middleware Logger será executado para todas as solicitações recebidas pelo aplicativo, antes que o manipulador de solicitação correspondente seja chamado.

O middleware Logger registra informações sobre cada solicitação, incluindo o método HTTP, a rota solicitada, o endereço IP do cliente e o tempo necessário para lidar com a solicitação. Isso pode ser útil para fins de depuração e monitoramento de desempenho do aplicativo.

Além do middleware Logger, a [https://github.com/gojeffotoni/quick](https://github.com/gojeffotoni/quick) inclui vários outros exemplos de middleware, incluindo o middleware Recovery para lidar com erros panics, o middleware CORS para habilitar o acesso cruzado a recursos e o middleware JWT para autenticação baseada em tokens JWT.
