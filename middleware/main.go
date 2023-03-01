package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gojeffotoni/quick"
	"github.com/gojeffotoni/quick/internal/concat"
	"github.com/gojeffotoni/quick/middleware/cors"
	"github.com/gojeffotoni/quick/middleware/msgid"
)

func main() {
	app := quick.New()
	app.Use(cors.New(), "cors")
	println("..................")

	group := app.Group("/v1")

	group.Post("/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		type My struct {
			Name string `json:"name"`
			Year int    `json:"year"`
		}

		var my My
		err := c.BodyParser(&my)
		fmt.Println("byte:", c.Body())

		if err != nil {
			c.Status(400).SendString(err.Error())
			return
		}

		fmt.Println("String:", c.BodyString())
		c.Status(200).JSON(&my)
		return
	})

	app.Post("/v2/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		type My struct {
			Name string `json:"name"`
			Year int    `json:"year"`
		}

		var my2 My
		err := c.Bind(&my2)
		if err != nil {
			c.Status(400).SendString(err.Error())
			return
		}

		// b, err := json.Marshal(&my2)
		// fmt.Println("String:", c.BodyString())

		c.Status(200).JSON(&my2)
		return
	})

	// app.Use(func(h http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		//This is the first middleware! This guy generate a msgID
	// 		r.Header.Set("Messageid", "12345")
	// 		w.Header().Set("Messageid", "12345")
	// 		h.ServeHTTP(w, r)
	// 	})
	// })

	// app.Use(msgid.New())

	// app.Use(msgid.New(msgid.Config{
	// 	Name: "Myid",
	// }))

	// exemplo msgid
	// app.Use(msgid.New(msgid.Config{
	// 	Name:  "Myid",
	// 	Start: 50000,
	// 	End:   100000,
	// }))

	// ex msgi MIDDLEWARE
	app.Use(msgid.New(msgid.Config{
		Name: "Myid",
		Algo: func() string {
			rand.Seed(time.Now().UnixNano())
			return concat.String("TID", strconv.Itoa(100000000+int(rand.Intn(100000000))))
		},
	}))

	app.Get("/v1/customer/:param1/:param2", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")

		type my struct {
			Msg string `json:"msg"`
			Key string `json:"key"`
			Val string `json:"val"`
		}

		log.Println(c.Headers["Messageid"])

		c.Status(200).JSON(&my{
			Msg: "Quick ❤️",
			Key: c.Param("param1"),
			Val: c.Param("param2"),
		})
	})

	app.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//This is the third middleware! This guy blocks request if header contains Block == true
			if r.Header.Get("Block") == "" {
				w.WriteHeader(400)
				w.Write([]byte(`"Message": "send the block header, please! :("`))
				return
			}

			if r.Header.Get("Block") == "true" {
				w.WriteHeader(200)
				w.Write([]byte(`"Message": "quicks block this request :)"`))
				return
			}
			h.ServeHTTP(w, r)
		})
	})

	app.Get("/v1/blocked", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")

		type my struct {
			Msg   string `json:"msg"`
			Block string `json:"block_message"`
		}

		log.Println(c.Headers["Messageid"])

		c.Status(200).JSON(&my{
			Msg:   "Quick ❤️",
			Block: c.Headers["Block"][0],
		})
	})

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
