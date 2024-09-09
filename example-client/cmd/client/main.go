package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	httpClient *http.Client
}

var cli *Client

// init and teardown functions

func init() {
	cli = &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func teardown() {
	// fmt.Println("Tearing down client...")
	cli.httpClient.CloseIdleConnections()
	// fmt.Println("Graceful shutdown complete.")
}

// main function

func main() {
	defer teardown()

	var hostname = make(chan string)
	var uniqueHostnamesRequestCount = make(map[string]int)
	var wg sync.WaitGroup

	go func() {
		for msg := range hostname {
			uniqueHostnamesRequestCount[msg]++
		}
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendRequest(hostname)
		}()
	}

	wg.Wait()
	close(hostname)

	fmt.Println("Unique hostnames:", uniqueHostnamesRequestCount)
}

// concurrency functions
func sendRequest(c chan<- string) {
	msg, err := cli.getBackend()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	c <- msg
	// fmt.Println("Backend message:", msg)
}

// Client methods

func (c *Client) getBackend() (string, error) {
	res, err := c.httpClient.Get("http://localhost:80")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var backendResponse BackendResponse
	err = json.Unmarshal(resBody, &backendResponse)
	if err != nil {
		return "", err
	}

	return backendResponse.Hostname, nil
}

type BackendResponse struct {
	Message        string `json:"message"`
	CurrentTime    string `json:"current_time"`
	ServiceAddress string `json:"service_address"`
	Hostname       string `json:"hostname"`
}
