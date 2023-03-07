package main

import (
	"github.com/jeffotoni/quick"
)

func main() {
	app := quick.New()
 
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

		 app.Get("/greet/:name", func(c *quick.Ctx) error {
		name := c.Param("name")
		c.Set("Content-Type", "text/plain")
		return c.Status(200).SendString("Olá " + name + "!")
	})

	app.Listen("0.0.0.0:8080")
}