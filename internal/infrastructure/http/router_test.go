package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t4vr1s/go-service-template/internal/usecase"
)

func TestNewRouterHealthEndpoint(t *testing.T) {
	router := NewRouter(usecase.HealthCheck{}, usecase.Greeter{})
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unexpected response body: %v", err)
	}

	if response["status"] != "ok" {
		t.Fatalf("expected status ok, got %q", response["status"])
	}
}

func TestNewRouterGreetEndpointUsesDefaultName(t *testing.T) {
	router := NewRouter(usecase.HealthCheck{}, usecase.Greeter{})
	request := httptest.NewRequest(http.MethodGet, "/greet", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unexpected response body: %v", err)
	}

	if response["message"] != "hello, world" {
		t.Fatalf("expected default greeting, got %q", response["message"])
	}
}

func TestNewRouterGreetEndpointUsesProvidedName(t *testing.T) {
	router := NewRouter(usecase.HealthCheck{}, usecase.Greeter{})
	request := httptest.NewRequest(http.MethodGet, "/greet?name=travis", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("unexpected response body: %v", err)
	}

	if response["message"] != "hello, travis" {
		t.Fatalf("expected named greeting, got %q", response["message"])
	}
}

func TestNewRouterReturnsNotFoundForUnknownPath(t *testing.T) {
	router := NewRouter(usecase.HealthCheck{}, usecase.Greeter{})
	request := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}
