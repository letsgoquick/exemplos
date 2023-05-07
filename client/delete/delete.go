package main

import (
	"log"

	"github.com/jeffotoni/quick/httpclient"
)

func main() {
	resp := httpclient.Delete("http://localhost:8000/delete")
	log.Printf("response: %s", resp.Body)
}
