package main

import "github.com/gojeffotoni/quick"
       "net/http"

func main() {
app.Use(func(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      //This is the third middleware! This guy blocks request if header contains Block == true
      if r.Header.Get("Block") == "" {
        w.WriteHeader(400)
        w.Write([]byte("Message": "send the block header, please! :("))
        return
      }

      if r.Header.Get("Block") == "true" {
        w.WriteHeader(200)
        w.Write([]byte("Message": "quicks block this request :)"))
        return
      }
      h.ServeHTTP(w, r)
    })
  })

  app.Get("/v1/blocked", func(c *quick.Ctx) {
    c.Set("Content-Type", "application/json")

    type my struct {
      Msg   string json:"msg"
      Block string json:"block_message"
    }

    log.Println(c.Headers["Messageid"])

    c.Status(200).JSON(&my{
      Msg:   "Quick ❤️",
      Block: c.Headers["Block"][0],
    })
  })

  log.Fatal(app.Listen("0.0.0.0:8080"))

}