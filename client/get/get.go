package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jeffotoni/quick/httpclient"
)

func main() {
	callLocally()
	// callLetsGoQuick()
	// callWithCustomClient()
}

func callLocally() {
	resp, err := httpclient.Get("http://localhost:8000/get")
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}

func callLetsGoQuick() {
	resp, err := httpclient.Get("https://letsgoquick.com")
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}

func callGoDev() {
	resp, err := httpclient.Delete("https://go.dev")
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

	resp, err := c.Delete("http://localhost:8000/get")
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Printf("response: %s | statusCode: %d", resp.Body, resp.StatusCode)
}
