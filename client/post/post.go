package main

import (
	"io"
	"log"
	"strings"

	"github.com/jeffotoni/quick/httpclient"
)

func main() {
	resp := httpclient.Post("http://localhost:8000/post", io.NopCloser(strings.NewReader(`{"data": "client quick!"}`)))
	log.Printf("response: %s", resp.Body)
}
