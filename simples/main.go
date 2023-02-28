package main

import (
	"fmt"
	"net/http"

	"github.com/gojeffotoni/quick"
	"github.com/gojeffotoni/quick/middleware/cors"
)

func main() {
	app := quick.New(quick.Config{
		MaxBodySize: 5 * 1024 * 1024,
	})

	type My struct {
		Name string `json:"name"`
		Year int    `json:"year"`
	}
	app.Use(cors.New())

	// group := app.Group("/v3")
	// group.Get("/user", func(c *quick.Ctx) {
	// 	c.Status(200).SendString("GET guser ok!!!")
	// 	return
	// })
	// group.Post("/user", func(c *quick.Ctx) {
	// 	c.Status(200).SendString("[POST] guser ok!!!")
	// 	return
	// })

	app.Get("/v2/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação com Cors❤️!")
	})

	app.Post("/v1/user", func(c *quick.Ctx) {
		var my My
		err := c.BodyParse(&my)
		if err != nil {
			fmt.Println("error Body:", err)
			c.Status(http.StatusBadRequest)
			return
		}
		c.Status(200).JSON(&my)
		return
	})

	app.Get("/v1/customer/:param1/:param2", func(c *quick.Ctx) {
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
		return
	})

	app.Get("/v1/user", func(c *quick.Ctx) {
		c.Set("Content-Type", "application/json")
		c.Status(200).SendString("Quick em ação ❤️!")
		return
	})

	for k, v := range app.GetRoute() {
		fmt.Println(k, "[", v, "]")
	}
	app.Listen("0.0.0.0:8080")
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	//mw "github.com/lovquick/middleware"
// 	"github.com/lovquick/quick"
// )

// func MyQuick1(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(quick.ConcatStr("Quick => ", r.RequestURI, " ...")))
// }

// func MyQuick2(c *quick.Ctx) {
// 	println("Params")
// 	for k, v := range c.Params {
// 		fmt.Println("p:", k, " = ", v)
// 	}

// 	println("Query")
// 	for k, v := range c.Query {
// 		fmt.Println("q:", k, " = ", v)
// 	}

// 	println("Headers")
// 	for k, v := range c.Headers {
// 		fmt.Println("h:", k, " = ", v)
// 	}

// 	c.Response.Header().Add("Content-Type", "application/json")
// 	c.Response.WriteHeader(http.StatusOK)
// 	c.Response.Write([]byte(quick.ConcatStr("Quick => ", c.Request.RequestURI, " ...")))
// }

// type MyStruct struct {
// 	Name string `json:"name"`
// 	Year int    `json:"year"`
// }

// type Out struct {
// 	Error string `json:"error,omitempty"`
// 	Msg   string `json:"msg"`
// }

// func MyQuick3(c *quick.Ctx) {

// 	for k, v := range c.Headers {
// 		fmt.Println("h:", k, " = ", v)
// 	}

// 	var my MyStruct

// 	err := c.BodyParse(&my)
// 	if err != nil {
// 		fmt.Println("error Body:", err)
// 		c.Response.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("json Struct:", c.BodyString())
// 	fmt.Println("my struct:", my)
// 	fmt.Println("my struct:", my.Name, " - Year: ", my.Year)

// 	c.Set("Content-Type", "application/json")
// 	c.Set("Generate", "Quick")
// 	c.Set("AppID", "T8201992X39393")

// 	if my.Year < 2020 {

// 		var out = Out{
// 			Error: "Error here",
// 			Msg:   "somente um test",
// 		}

// 		c.Status(400).JSON(out)
// 		return
// 	}

// 	if my.Year < 2022 {
// 		bb := []byte(`{"msg":"success to 2022"}`)
// 		c.Status(400).Byte(bb)
// 		return
// 	}

// 	c.Status(200).JSON(Out{
// 		Msg: "ok",
// 	})
// }

// func main() {

// 	q := quick.New()

// 	q.Get("/v1/user2/:param1/:param2/:param3", MyQuick2)

// 	q.Get("/v1/user", MyQuick2)

// 	// q.Use(mw.Logger)
// 	q.Post("/v1/user", MyQuick3)

// 	for k, v := range q.ListRoutes() {
// 		fmt.Println(k, " = ", v.Method, " = ", v.Path)
// 	}

// 	q.Listen("0.0.0.0:8080")

// }
