package main

import (
	"log"

	"github.com/jeffotoni/quick/httpclient"
)

func main() {
	resp := httpclient.Get("http://localhost:8000/get")

	log.Printf("response: %s", resp.Body)
}
