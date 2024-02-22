package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// NewBackoffClient returns a new configured http client with request retrying and backoff.
func NewBackoffClient(requestTimeout time.Duration, maxRetries int) *http.Client {
	// Create an HTTP client
	client := &http.Client{}

	// Configure retrying and backoff
	client.Transport = backoff.Transport{
		// Wrap the default transport
		Delegate: http.DefaultTransport,

		// Configure exponential backoff
		Backoff: backoff.NewExponentialBackOff(),

		// Maximum number of retries
		MaxRetries: uint64(maxRetries),
	}

	// Set timeout for the entire request including retries
	client.Timeout = requestTimeout

	return client
}

func main() {
	// Create a new HTTP client with retrying and backoff
	client := NewBackoffClient(5*time.Second, 3) // 5 seconds timeout, maximum 3 retries

	// Send a GET request to example.com
	resp, err := client.Get("http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Print response status code and body
	fmt.Println("Status Code:", resp.StatusCode)
	// Read and print the response body if needed
}
