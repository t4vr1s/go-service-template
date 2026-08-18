package main

import (
	"log"

	"github.com/t4vr1s/go-service-template/internal/infrastructure/config"
	httpserver "github.com/t4vr1s/go-service-template/internal/infrastructure/http"
)

func main() {
	cfg := config.Load()
	server := httpserver.NewServer(cfg)

	log.Printf("starting HTTP server on %s", cfg.Address())
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
