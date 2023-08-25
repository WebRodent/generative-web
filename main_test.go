package main

// tests for main.go
import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"generative-web/internal/config"
)

func TestServer(t *testing.T) {
	conf, err := config.Load("config.yml")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	var endpoints = []string{
		"/",
		"/ping",
		"/status",
		"/template-load",
		"/template-load?template=single",
	}
	var wg sync.WaitGroup
	go Start(&wg)
	// halt program to allow server to start
	time.Sleep(1 * time.Second)
	// make a ping request to the server
	var baseUrl string = fmt.Sprintf("http://%s:%d", conf.Server.Host, conf.Server.Port)
	for _, endpoint := range endpoints {
		response, err := http.Get(fmt.Sprintf("%s%s", baseUrl, endpoint))
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if response.StatusCode != 200 {
			t.Errorf("Expected 200, got %d", response.StatusCode)
		}
	}

	wg.Done()
}
