package http

import (
	"net/http"

	"github.com/t4vr1s/go-service-template/internal/domain"
	"github.com/t4vr1s/go-service-template/pkg/httpx"
)

type HealthChecker interface {
	Execute() domain.HealthStatus
}

type Greeter interface {
	Execute(name string) domain.Greeting
}

func NewRouter(healthCheck HealthChecker, greeter Greeter) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]string{
			"service": "go-service-template",
			"status":  "running",
		})
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, healthCheck.Execute())
	})

	mux.HandleFunc("GET /greet", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, greeter.Execute(r.URL.Query().Get("name")))
	})

	return mux
}
