package http

import (
	"net/http"

	"github.com/t4vr1s/go-service-template/internal/infrastructure/config"
	"github.com/t4vr1s/go-service-template/internal/usecase"
)

func NewServer(cfg config.Config) *http.Server {
	return &http.Server{
		Addr:    cfg.Address(),
		Handler: NewRouter(usecase.HealthCheck{}, usecase.Greeter{}),
	}
}
