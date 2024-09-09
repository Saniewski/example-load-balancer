package main

import (
	"fmt"
	"net/http"

	"github.com/saniewski/test-swarm/example-go-service/internal/config"
	"github.com/saniewski/test-swarm/example-go-service/internal/routes"
)

func init() {
	config.Load()
}

func main() {
	cfg := config.Get()
	r := routes.Build()

	fmt.Printf("Starting server on %s:%s\n", cfg.Hostname, cfg.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Hostname, cfg.Port), r)
}
