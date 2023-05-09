package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/jeffotoni/quick/httpclient"
)

func main() {
	// callLocally()
	// callLetsGoQuick()
	// callGoDev()
	callWithCustomClient()
}

func callLocally() {
	resp, err := httpclient.Put("http://localhost:8000/put", io.NopCloser(strings.NewReader(`{"data": "client quick!"}`)))
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}

func callLetsGoQuick() {
	resp, err := httpclient.Put("https://letsgoquick.com", io.NopCloser(strings.NewReader(`{"data": "client quick!"}`)))
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}

func callGoDev() {
	resp, err := httpclient.Put("https://go.dev", io.NopCloser(strings.NewReader(`{"data": "client quick!"}`)))
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}

func callWithCustomClient() {
	c := httpclient.Client{
		Ctx: context.Background(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		ClientHttp: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    10,
				MaxConnsPerHost: 10,
			},
			Timeout: 0,
		},
	}

	resp, err := c.Put("http://localhost:8000/put", io.NopCloser(strings.NewReader(`{"data": "client quick!"}`)))
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}
