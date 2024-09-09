package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/saniewski/test-swarm/example-go-service/internal/config"
)

type jsonResponse struct {
	Message        string `json:"message"`
	CurrentTime    string `json:"current_time"`
	ServiceAddress string `json:"service_address"`
	Hostname       string `json:"hostname"`
}

func Get(w http.ResponseWriter, r *http.Request) {

	cfg := config.Get()
	currentTime := time.Now().UTC().String()
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := jsonResponse{
		Message:        "Hello, Example Go Service!",
		CurrentTime:    currentTime,
		ServiceAddress: fmt.Sprintf("%s:%s", cfg.Hostname, cfg.Port),
		Hostname:       hostname,
	}

	out, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
