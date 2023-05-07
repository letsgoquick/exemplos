package main

import (
	"io"
	"log"
	"strings"

	"github.com/jeffotoni/quick/httpclient"
)

func main() {
	resp := httpclient.Put("http://localhost:8000/put", io.NopCloser(strings.NewReader(`{"data": "client quick!"}`)))
	log.Printf("response: %s", resp.Body)
}
