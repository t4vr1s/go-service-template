package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/t4vr1s/go-service-template/internal/infrastructure/config"
	httpserver "github.com/t4vr1s/go-service-template/internal/infrastructure/http"
)

func main() {
	cfg := config.Load()
	server := httpserver.NewServer(cfg)

	log.Printf("starting HTTP server on %s", cfg.Address())
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
