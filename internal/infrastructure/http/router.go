package http

import (
	"net/http"

	"github.com/t4vr1s/go-service-template/internal/usecase"
	"github.com/t4vr1s/go-service-template/pkg/httpx"
)

func NewRouter(healthCheck usecase.HealthCheck, greeter usecase.Greeter) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]string{
			"service": "go-service-template",
			"status":  "running",
		})
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, healthCheck.Execute())
	})

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, greeter.Execute(r.URL.Query().Get("name")))
	})

	return mux
}
